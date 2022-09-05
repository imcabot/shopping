package api

import (
	"github.com/gin-gonic/gin"
	"log"
	cartApi "shoping/api/cart"
	categoryApi "shoping/api/category"
	orderApi "shoping/api/order"
	productApi "shoping/api/product"
	userApi "shoping/api/user"
	"shoping/category"
	"shoping/config"
	"shoping/domain/cart"
	"shoping/domain/order"
	"shoping/domain/product"
	"shoping/domain/user"
	"shoping/utils/database_handler"
	"shoping/utils/middleware"
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
	db := database_handler.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)
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
	categoryGroup.POST("/upload", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey),
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
	cartGroup := r.Group("/cart", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey))
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
	orderConTroller := orderApi.NewOrderController(orderService)
	orderGroup := r.Group("/order")
	orderGroup.POST("", orderConTroller.CompleteOrder)
	orderGroup.DELETE("", orderConTroller.CancelOrder)
	orderGroup.GET("", orderConTroller.GetOrders)

}
