package dto

type AfterSaleInfo struct {
	ProductName   string  `form:"productName,required" json:"productName,required" query:"productName,required"`
	Files         []*File `form:"files,omitempty" json:"files,omitempty" query:"files,omitempty"`
	AfterSaleText string  `form:"afterSaleText,required" json:"afterSaleText,required" query:"afterSaleText,required"`
}

type CheckPurchasedReq struct {
	UserId    int64 `form:"userId,required" json:"userId,required" query:"userId,required"`
	ProductId int64 `form:"productId,required" json:"productId,required" query:"productId,required"`
}

type CheckPurchasedResp struct {
	AfterSaleInfo *AfterSaleInfo `form:"afterSaleInfo,omitempty" json:"afterSaleInfo,omitempty" query:"afterSaleInfo,omitempty"`
	Code          int64          `form:"code,required" json:"code,required" query:"code,required"`
	Message       string         `form:"message,required" json:"message,required" query:"message,required"`
}

type EduProduct struct {
	ProductId   int64  `form:"productId,required" json:"productId,required" query:"productId,required"`
	ProductName string `form:"productName,required" json:"productName,required" query:"productName,required"`
	ProductDesc string `form:"productDesc,required" json:"productDesc,required" query:"productDesc,required"`
	ProductImg  string `form:"productImg,required" json:"productImg,required" query:"productImg,required"`
	ProductType int32  `form:"productType,required" json:"productType,required" query:"productType,required"`
	CreatedTime int64  `form:"createdTime,required" json:"createdTime,required" query:"createdTime,required"`
	SaleVolume  int32  `form:"saleVolume,required" json:"saleVolume,required" query:"saleVolume,required"`
	Inventory   int32  `form:"inventory,required" json:"inventory,required" query:"inventory,required"`
	Provider    string `form:"provider,required" json:"provider,required" query:"provider,required"`
}

type File struct {
	FileType int32  `form:"fileType,required" json:"fileType,required" query:"fileType,required"`
	FileName string `form:"fileName,required" json:"fileName,required" query:"fileName,required"`
	FileUrl  string `form:"fileUrl,required" json:"fileUrl,required" query:"fileUrl,required"`
	FileId   int64  `form:"fileId,required" json:"fileId,required" query:"fileId,required"`
}

type GetProductDetailReq struct {
	ProductId int64 `form:"productId,required" json:"productId,required" query:"productId,required"`
}

type GetProductDetailResp struct {
	Detail  *ProductDetail `form:"detail,omitempty" json:"detail,omitempty" query:"detail,omitempty"`
	Code    int64          `form:"code,required" json:"code,required" query:"code,required"`
	Message string         `form:"message,required" json:"message,required" query:"message,required"`
}

type GetProductEditInfoReq struct {
	ProductId int64 `form:"productId,required" json:"productId,required" query:"productId,required"`
}

type GetProductEditInfoResp struct {
	Info    *ProductEditInfo `form:"info,omitempty" json:"info,omitempty" query:"info,omitempty"`
	Code    int64            `form:"code,required" json:"code,required" query:"code,required"`
	Message string           `form:"message,required" json:"message,required" query:"message,required"`
}

type GetPurchaseRecordsReq struct {
	UserId      int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	ProductName string `form:"productName,required" json:"productName,required" query:"productName,required"`
	PageNum     int32  `form:"pageNum,required" json:"pageNum,required" query:"pageNum,required"`
	PageSize    int32  `form:"pageSize,required" json:"pageSize,required" query:"pageSize,required"`
}

type GetPurchaseRecordsResp struct {
	PurchaseRecords []*PurchaseRecord `form:"PurchaseRecords,required" json:"PurchaseRecords,required" query:"PurchaseRecords,required"`
	Code            int64             `form:"code,required" json:"code,required" query:"code,required"`
	Message         string            `form:"message,required" json:"message,required" query:"message,required"`
}

type GetUserInfoReq struct {
	UserId int64 `form:"userId,required" json:"userId,required" query:"userId,required"`
}

type GetUserInfoResp struct {
	UserInfo *UserInfo `form:"userInfo,omitempty" json:"userInfo,omitempty" query:"userInfo,omitempty"`
	Code     int64     `form:"code,required" json:"code,required" query:"code,required"`
	Message  string    `form:"message,required" json:"message,required" query:"message,required"`
}

type LoginRequest struct {
	Account  string `form:"account,required" json:"account,required" query:"account,required"`
	Password string `form:"password,required" json:"password,required" query:"password,required"`
	UserType int32  `form:"userType,required" json:"userType,required" query:"userType,required"`
}

type LoginResponse struct {
	UserId  int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	Code    int64  `form:"code,required" json:"code,required" query:"code,required"`
	Message string `form:"message,required" json:"message,required" query:"message,required"`
}

type ProductDetail struct {
	Id          int64    `form:"id,required" json:"id,required" query:"id,required"`
	ProductName string   `form:"productName,required" json:"productName,required" query:"productName,required"`
	ProductType int32    `form:"productType,required" json:"productType,required" query:"productType,required"`
	Price       int32    `form:"price,required" json:"price,required" query:"price,required"`
	Status      int32    `form:"status,required" json:"status,required" query:"status,required"`
	ProductDesc string   `form:"productDesc,required" json:"productDesc,required" query:"productDesc,required"`
	BannerImgs  []string `form:"bannerImgs,required" json:"bannerImgs,required" query:"bannerImgs,required"`
	Inventory   int32    `form:"inventory,required" json:"inventory,required" query:"inventory,required"`
	SaleVolume  int32    `form:"saleVolume,required" json:"saleVolume,required" query:"saleVolume,required"`
	Files       []*File  `form:"files,omitempty" json:"files,omitempty" query:"files,omitempty"`
}

type ProductEditInfo struct {
	Id            int64    `form:"id,required" json:"id,required" query:"id,required"`
	ProductName   string   `form:"productName,required" json:"productName,required" query:"productName,required"`
	ProductType   int32    `form:"productType,required" json:"productType,required" query:"productType,required"`
	Price         int32    `form:"price,required" json:"price,required" query:"price,required"`
	Files         []*File  `form:"files,omitempty" json:"files,omitempty" query:"files,omitempty"`
	Status        int32    `form:"status,required" json:"status,required" query:"status,required"`
	ProductDesc   string   `form:"productDesc,required" json:"productDesc,required" query:"productDesc,required"`
	BannerImgs    []string `form:"bannerImgs,required" json:"bannerImgs,required" query:"bannerImgs,required"`
	AfterSaleText string   `form:"afterSaleText,required" json:"afterSaleText,required" query:"afterSaleText,required"`
	Inventory     int32    `form:"inventory,required" json:"inventory,required" query:"inventory,required"`
}

type PurchaseRecord struct {
	RecordId    int64    `form:"recordId,required" json:"recordId,required" query:"recordId,required"`
	ProductId   int64    `form:"productId,required" json:"productId,required" query:"productId,required"`
	ProductName string   `form:"productName,required" json:"productName,required" query:"productName,required"`
	CreatedTime int64    `form:"createdTime,required" json:"createdTime,required" query:"createdTime,required"`
	ProductDesc string   `form:"productDesc,required" json:"productDesc,required" query:"productDesc,required"`
	DescImg     []string `form:"descImg,required" json:"descImg,required" query:"descImg,required"`
	ProductType int32    `form:"productType,required" json:"productType,required" query:"productType,required"`
	Payment     int32    `form:"payment,required" json:"payment,required" query:"payment,required"`
}

type PurchaseReq struct {
	UserId    int64 `form:"userId,required" json:"userId,required" query:"userId,required"`
	ProductId int64 `form:"productId,required" json:"productId,required" query:"productId,required"`
}

type PurchaseResp struct {
	Code    int64  `form:"code,required" json:"code,required" query:"code,required"`
	Message string `form:"message,required" json:"message,required" query:"message,required"`
}

type RechargeReq struct {
	UserId int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	CDkey  string `form:"CDkey,required" json:"CDkey,required" query:"CDkey,required"`
}

type RechargeResp struct {
	Code    int64  `form:"code,required" json:"code,required" query:"code,required"`
	Message string `form:"message,required" json:"message,required" query:"message,required"`
}

type RegisterRequest struct {
	Account  string `form:"account,required" json:"account,required" query:"account,required"`
	Password string `form:"password,required" json:"password,required" query:"password,required"`
	Nickname string `form:"nickname,required" json:"nickname,required" query:"nickname,required"`
}

type RegisterResponse struct {
	UserId  int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	Code    int64  `form:"code,required" json:"code,required" query:"code,required"`
	Message string `form:"message,required" json:"message,required" query:"message,required"`
}

type SearchEduProductsReq struct {
	Provider    string `form:"provider,required" json:"provider,required" query:"provider,required"`
	Keywords    string `form:"keywords,required" json:"keywords,required" query:"keywords,required"`
	ProductType int32  `form:"productType,required" json:"productType,required" query:"productType,required"`
	PageNum     int32  `form:"pageNum,required" json:"pageNum,required" query:"pageNum,required"`
	PageSize    int32  `form:"pageSize,required" json:"pageSize,required" query:"pageSize,required"`
}

type SearchEduProductsResp struct {
	EduProductList []*EduProduct `form:"EduProductList,required" json:"EduProductList,required" query:"EduProductList,required"`
	Code           int64         `form:"code,required" json:"code,required" query:"code,required"`
	Message        string        `form:"message,required" json:"message,required" query:"message,required"`
}

type UpsertEduProductReq struct {
	Id            int64    `form:"id,required" json:"id,required" query:"id,required"`
	UserId        int64    `form:"userId,required" json:"userId,required" query:"userId,required"`
	ProductName   string   `form:"productName,required" json:"productName,required" query:"productName,required"`
	ProductType   int32    `form:"productType,required" json:"productType,required" query:"productType,required"`
	Price         int32    `form:"price,required" json:"price,required" query:"price,required"`
	FileType      int32    `form:"fileType,required" json:"fileType,required" query:"fileType,required"`
	FileUrl       string   `form:"fileUrl,required" json:"fileUrl,required" query:"fileUrl,required"`
	Status        int32    `form:"status,required" json:"status,required" query:"status,required"`
	ProductDesc   string   `form:"productDesc,required" json:"productDesc,required" query:"productDesc,required"`
	BannerImgs    []string `form:"bannerImgs,required" json:"bannerImgs,required" query:"bannerImgs,required"`
	AfterSaleText string   `form:"afterSaleText,required" json:"afterSaleText,required" query:"afterSaleText,required"`
	FileName      string   `form:"fileName,required" json:"fileName,required" query:"fileName,required"`
	FileId        int64    `form:"fileId,required" json:"fileId,required" query:"fileId,required"`
}

type UpsertEduProductResp struct {
	Code    int64  `form:"code,required" json:"code,required" query:"code,required"`
	Message string `form:"message,required" json:"message,required" query:"message,required"`
}

type UserInfo struct {
	UserId   int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	UserType int32  `form:"userType,required" json:"userType,required" query:"userType,required"`
	Nickname string `form:"nickname,required" json:"nickname,required" query:"nickname,required"`
	Balance  int32  `form:"balance,required" json:"balance,required" query:"balance,required"`
	Account  string `form:"account,required" json:"account,required" query:"account,required"`
}
