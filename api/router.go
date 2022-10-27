package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"log"
	cartApi "shopping/api/cart"
	categoryApi "shopping/api/category"
	orderApi "shopping/api/order"
	productApi "shopping/api/product"
	userApi "shopping/api/user"
	"shopping/category"
	"shopping/config"
	"shopping/domain/cart"
	"shopping/domain/order"
	"shopping/domain/product"
	"shopping/domain/user"
	"shopping/utils/database_handler"
	"shopping/utils/middleware"
)

type Databases struct {
	categoryRepository  *category.Repository
	userRepository      *user.Repository
	productRepository   *product.Repository
	cartRepository      *cart.Repository
	cartItemRepository  *cart.ItemRepository
	orderRepository     *order.Repository
	orderItemRepository *order.OrderedItemRepository
}

var AppConfig = &config.ConfigUration{}

//根据配置文件创建数据库
func CreatDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		return nil
	}
	if err != nil {
		log.Fatalf("读取配置文件失败。 %v", err.Error())
	}
	dns := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.DatabaseSettings.Username, AppConfig.DatabaseSettings.Password,
		AppConfig.DatabaseSettings.Host, AppConfig.DatabaseSettings.Port, AppConfig.DatabaseSettings.DatabaseName)
	db := database_handler.NewMySQLDB(dns)
	return &Databases{
		categoryRepository:  category.NewCategoryRepository(db),
		userRepository:      user.NewUserRepository(db),
		productRepository:   product.NewProductRepository(db),
		cartRepository:      cart.NewCartRepository(db),
		cartItemRepository:  cart.NewItemRepository(db),
		orderRepository:     order.NewOrderRepository(db),
		orderItemRepository: order.NewOrderedItemRepository(db),
	}
}

//注册所有控制器
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreatDBs()
	RegisterUserHandlers(r, dbs)
	RegisterCategoryHandlers(r, dbs)
	RegisterCartHandlers(r, dbs)
	RegisterProductHandlers(r, dbs)
	RegisterOrderHandlers(r, dbs)

}

//注册分类控制器
func RegisterCategoryHandlers(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")
	categoryGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateCategory)
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST("/upload", middleware.AuthAdminMiddleware(AppConfig.SecretKey),
		categoryController.BulkCreateCategory)

}

//注册用户控制器
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/login", userController.Login)

}

//注册购物车控制器
func RegisterCartHandlers(r *gin.Engine, dbs Databases) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart", middleware.AuthUserMiddleware(AppConfig.SecretKey))
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/item", cartController.GetCart)

}

//注册商品控制器
func RegisterProductHandlers(r *gin.Engine, dbs Databases) {
	productService := product.NewService(*dbs.productRepository)
	productController := productApi.NewProductController(productService)
	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST("", middleware.AuthAdminMiddleware(AppConfig.SecretKey), productController.CreateProduct)
	productGroup.DELETE("", middleware.AuthAdminMiddleware(AppConfig.SecretKey), productController.DeleteProduct)
	productGroup.PATCH("", middleware.AuthAdminMiddleware(AppConfig.SecretKey), productController.UpdateProduct)
}

//注册订单控制器
func RegisterOrderHandlers(r *gin.Engine, dbs Databases) {
	orderService := order.NewService(*dbs.orderRepository, *dbs.orderItemRepository,
		*dbs.productRepository, *dbs.cartRepository, *dbs.cartItemRepository)
	orderController := orderApi.NewOrderController(orderService)
	orderGroup := r.Group("/order", middleware.AuthUserMiddleware(AppConfig.SecretKey))
	orderGroup.POST("", orderController.CompleteOrder)
	orderGroup.DELETE("", orderController.CancelOrder)
	orderGroup.GET("", orderController.GetOrders)

}

func CreatRDBs() {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		log.Fatalf("读取配置文件失败。 %v", err.Error())
	}
	addr := fmt.Sprintf("%v:%v", conf.RedisSettings.Host, conf.RedisSettings.Port)
	option := redis.Options{
		// Addr:     "localhost:6379", // windows
		Addr:     addr,
		Username: conf.RedisSettings.Username,
		Password: conf.RedisSettings.Password,     // no password set
		DB:       conf.RedisSettings.DatabaseName, // use default DB
	}
	database_handler.NewRedisDB(&option)

}
