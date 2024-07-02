package main

import (
	"Delimart/controllers"
	"Delimart/models"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	// // Set Access-Control-Allow-Origin to * for development
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// 		return next(c)
	// 	}
	// })

	var err error

	// Replace with your database credentials
	// dsn := "root:@tcp(127.0.0.1:3306)/db_ipat_uas_final"
	dsn := "delimart:8W@jQeSYDTPwpa!@tcp(mysql-delimart.alwaysdata.net:3306)/delimart_db_ipat_uas"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Verify database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Set database connection in models package
	models.SetDB(db)

	// Web controllers
	userController := &controllers.UserController{}
	barangController := &controllers.BarangController{}
	supplierController := &controllers.SupplierController{}
	pegawaiController := &controllers.PegawaiController{}
	pulsaController := &controllers.PulsaController{}
	roleController := &controllers.RoleController{}
	kategoriController := &controllers.KategoriController{}
	prefixController := &controllers.PrefixController{}
	providerController := &controllers.ProviderController{}
	LaporanDashboardController := &controllers.LaporanDashboardController{}

	// controllers dektop
	detailPenjualanController := &controllers.DetailPenjualanController{}
	detailPenerimaanController := &controllers.DetailPenerimaanController{}
	PenjualanController := &controllers.PenjualanController{}
	penerimaanController := &controllers.PenerimaanController{}

	// controllers mobile
	LaporanArusStokController := &controllers.LaporanArusStokController{}
	LaporanPenjualanBarangController := &controllers.LaporanPenjualanBarangController{}
	LaporanPenjualanHarianController := &controllers.LaporanPenjualanHarianController{}
	LaporanPenjualanBulananController := &controllers.LaporanPenjualanBulananController{}
	LaporanPenjualanTahunanController := &controllers.LaporanPenjualanTahunanController{}
	LaporanPenjualanAllController := &controllers.LaporanPenjualanAllController{}
	LaporanPenerimaanHarianController := &controllers.LaporanPenerimaanHarianController{}
	LaporanPenerimaanBulananController := &controllers.LaporanPenerimaanBulananController{}
	LaporanPenerimaanTahunanController := &controllers.LaporanPenerimaanTahunanController{}
	LaporanPenerimaanAllController := &controllers.LaporanPenerimaanAllController{}
	LaporanPenjualanHomeController := &controllers.LaporanPenjualanHomeController{}

	// auth controller
	authController := &controllers.AuthController{}

	// Routes
	// Welcome route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Go Echo Server!")
	})

	// auth routes
	e.POST("/login", authController.Login)
	e.POST("/logout/:kd_pegawai", authController.Logout)

	// User routes
	e.GET("/user", userController.GetAllUsers)
	e.GET("/user/:id", userController.GetUser)
	e.GET("/user/search", userController.SearchUser)
	e.POST("/user", userController.CreateUser)
	e.PUT("/user/:id", userController.UpdateUser)
	e.DELETE("/user/:id", userController.DeleteUser)

	// Barang routes
	e.GET("/barang", barangController.GetAllBarang)
	e.GET("/barang/:kd_barang", barangController.GetBarangByKodeBarang)
	e.GET("/barang/search", barangController.SearchBarang)
	e.POST("/barang", barangController.CreateBarang)
	e.PUT("/barang/:kd_barang", barangController.UpdateBarang)
	e.DELETE("/barang/:kd_barang", barangController.DeleteBarang)

	// Supplier routes
	e.GET("/supplier", supplierController.GetAllSuppliers)
	e.GET("/supplier/:kd_supplier", supplierController.GetSupplierByKodeSupplier)
	e.GET("/supplier/search", supplierController.SearchSupplier)
	e.POST("/supplier", supplierController.CreateSupplier)
	e.PUT("/supplier/:kd_supplier", supplierController.UpdateSupplier)
	e.DELETE("/supplier/:kd_supplier", supplierController.DeleteSupplier)

	// Pegawai routes
	e.GET("/pegawai", pegawaiController.GetAllPegawai)
	e.GET("/pegawai/:kd_pegawai", pegawaiController.GetPegawaiByKodePegawai)
	e.GET("/pegawai/search", pegawaiController.SearchPegawai)
	e.POST("/pegawai", pegawaiController.CreatePegawai)
	e.PUT("/pegawai/:kd_pegawai", pegawaiController.UpdatePegawai)
	e.DELETE("/pegawai/:kd_pegawai", pegawaiController.DeletePegawai)

	// Pulsa routes
	e.GET("/pulsa", pulsaController.GetAllPulsa)
	e.GET("/pulsa/:kd_pulsa", pulsaController.GetPulsaByKodePulsa)
	e.GET("/pulsa/search", pulsaController.SearchPulsa)
	e.POST("/pulsa", pulsaController.CreatePulsa)
	e.PUT("/pulsa/:kd_pulsa", pulsaController.UpdatePulsa)
	e.DELETE("/pulsa/:kd_pulsa", pulsaController.DeletePulsa)

	// Role routes
	e.GET("/role", roleController.GetAllRoles)
	e.GET("/role/:kd_role", roleController.GetRoleByKodeRole)
	e.POST("/role", roleController.CreateRole)
	e.PUT("/role/:kd_role", roleController.UpdateRole)
	e.DELETE("/role/:kd_role", roleController.DeleteRole)

	// Kategori routes
	e.GET("/kategori", kategoriController.GetAllKategori)
	e.GET("/kategori/:kd_kategori", kategoriController.GetKategoriByKodeKategori)
	e.GET("/kategori/search", kategoriController.SearchKategori)
	e.POST("/kategori", kategoriController.CreateKategori)
	e.PUT("/kategori/:kd_kategori", kategoriController.UpdateKategori)
	e.DELETE("/kategori/:kd_kategori", kategoriController.DeleteKategori)

	// Prefix routes
	e.GET("/prefix", prefixController.GetAllPrefix)
	e.GET("/prefix/:kd_prefix", prefixController.GetPrefixByKodePrefix)
	e.POST("/prefix", prefixController.CreatePrefix)
	e.PUT("/prefix/:kd_prefix", prefixController.UpdatePrefix)
	e.DELETE("/prefix/:kd_prefix", prefixController.DeletePrefix)

	// Provider routes
	e.GET("/provider", providerController.GetAllProviders)
	e.GET("/provider/:kd_provider", providerController.GetProviderByKodeProvider)
	e.POST("/provider", providerController.CreateProvider)
	e.PUT("/provider/:kd_provider", providerController.UpdateProvider)
	e.DELETE("/provider/:kd_provider", providerController.DeleteProvider)

	// routes desktop

	// barang
	e.GET("/get_barang1/:kd_supplier", barangController.GetBarangByKodeSupp)
	e.GET("/get_barang2/:nama", barangController.GetBarangByNama)

	// supplier
	e.GET("/get_supplier1/:nama", supplierController.GetKodeSupplierByNama)

	// pulsa
	e.GET("/get_pulsa/:provider", pulsaController.GetPulsaByProvider)

	// prefix
	e.GET("/get_provider/:prefix", prefixController.GetProviderByPrefix)

	// Detail Penjualan
	e.POST("/detail_penjualan", detailPenjualanController.CreateDetailPenjualan)

	// Detail Penerimaan
	e.POST("/detail_penerimaan", detailPenerimaanController.CreateDetailPenerimaan)

	// Penerimaan routes
	e.GET("/get_penerimaan/:nota", penerimaanController.GetPenerimaanByNota)
	e.GET("/nota_auto", penerimaanController.GetNotaAuto)
	// e.GET("/penerimaan/:kd_transaksi_terima", penerimaanController.GetPenerimaanByKodeTransaksiTerima)
	e.POST("/penerimaan", penerimaanController.CreatePenerimaan)
	e.PUT("/penerimaan/:kd_transaksi_terima", penerimaanController.UpdatePenerimaan)
	e.DELETE("/penerimaan/:kd_transaksi_terima", penerimaanController.DeletePenerimaan)
	e.DELETE("/delete_all_penerimaan/:no_terima", penerimaanController.DeleteAllPenerimaan)

	//  Penjualan
	e.GET("/struk_auto", PenjualanController.GetStrukAuto)
	e.GET("/get_penjualan/:struk", PenjualanController.GetData)
	e.POST("/penjualan", PenjualanController.CreatePenjualan)
	e.PUT("/penjualan/:kd_transaksi", PenjualanController.UpdatePenjualan)
	e.DELETE("/penjualan/:kd_transaksi", PenjualanController.DeletePenjualan)
	e.DELETE("/delete_all_penjualan/:struk", PenjualanController.DeleteAllPenjualan)

	// routes mobile
	// Get Laporan Arus Stok
	e.GET("/laporan_arus_stok", LaporanArusStokController.GetLaporanArusStok)

	// Get Laporan Penjualan Per barang
	e.GET("/laporan_penjualan_barang", LaporanPenjualanBarangController.GetLaporanPenjualanBarang)

	// Get Laporan Penjualan Harian dan Bulanan
	e.GET("/laporan_penjualan_harian", LaporanPenjualanHarianController.GetLaporanPenjualanHarian)
	e.GET("/laporan_penjualan_bulanan", LaporanPenjualanBulananController.GetLaporanPenjualanBulanan)
	e.GET("/laporan_penjualan_tahunan", LaporanPenjualanTahunanController.GetLaporanPenjualanTahunan)
	e.GET("/laporan_penjualan_all", LaporanPenjualanAllController.GetLaporanPenjualanAll)

	// Get Laporan Penerimaan Harian dan Bulanan
	e.GET("/laporan_penerimaan_harian", LaporanPenerimaanHarianController.GetLaporanPenerimaanHarian)
	e.GET("/laporan_penerimaan_bulanan", LaporanPenerimaanBulananController.GetLaporanPenerimaanBulanan)
	e.GET("/laporan_penerimaan_tahunan", LaporanPenerimaanTahunanController.GetLaporanPenerimaanTahunan)
	e.GET("/laporan_penerimaan_all", LaporanPenerimaanAllController.GetLaporanPenerimaanAll)

	// get laporan menu home mobile
	e.GET("/laporan_penjualan_home", LaporanPenjualanHomeController.GetLaporanPenjualanHome)
	// get laporan menu dashboard web
	e.GET("/dashboard", LaporanDashboardController.GetLaporanDashboard)
	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
