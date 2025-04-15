package router

import (
	"os"
	"singo/controller"
	"singo/controller/auth"
	"singo/controller/customer"
	"singo/controller/platform"
	"singo/controller/supplier"
	"singo/controller/user"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)

		// 用户登录
		v1.POST("user/register", auth.UserRegister)

		// 用户登录
		v1.POST("user/login", auth.UserLogin)

		// 需要登录保护的
		_auth := v1.Group("")
		_auth.Use(middleware.AuthRequired())
		{
			// User Routing
			_auth.GET("user/me", auth.UserMe)
			_auth.DELETE("user/logout", auth.UserLogout)

			// 平台管理端
			_platform := _auth.Group("/platform")
			_platform.POST("/get_supplier_list", platform.GetSupplierList)  // 获取供应商列表
			_platform.POST("/get_customer_list", platform.GetCustomerList)  // 获取业务方列表
			_platform.POST("/get_user_list", platform.GetUserList)          // 获取用户列表
			_platform.POST("/get_coupon_stat", platform.GetCouponStat)      // 获取发券数据
			_platform.POST("/get_write_off_stat", platform.GetWriteOffStat) // 获取核销数据

			// 服务商端
			_customer := _auth.Group("/customer")
			_customer.POST("/get_order_list", customer.GetOrderList)   // 获取订单列表
			_customer.POST("/get_coupon_stat", customer.GetCouponStat) // 获取发券数据
			_customer.POST("/send_coupon", customer.SendCoupon)        // 发券

			// 业务方端
			_supplier := _auth.Group("/supplier")
			_supplier.POST("get_write_off_stat", supplier.GetWriteOffStat) // 获取核销记录
			_supplier.POST("check_coupon", supplier.CheckCoupon)           // 核销

			// 用户端
			_user := _auth.Group("/user")
			_user.POST("get_coupon_list", user.GetCouponList) // 获取码券列表
			_user.POST("apply_coupon", user.ApplyCoupon)      // 申请获取码券
		}
	}
	return r
}
