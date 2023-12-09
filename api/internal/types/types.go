// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// required : true
	// min : 0
	Page uint64 `json:"page" validate:"required,number,gt=0"`
	// Page size | 单页数据行数
	// required : true
	// max : 100000
	PageSize uint64 `json:"pageSize" validate:"required,number,lt=100000"`
}

// The page request parameters | 列表请求参数
// swagger:model PageAtPathInfo
type PageAtPathInfo struct {
	// Page number | 第几页
	// required : true
	// min : 0
	Page uint64 `form:"page" validate:"required,number,gt=0"`
	// Page size | 单页数据行数
	// required : true
	// max : 100000
	PageSize uint64 `form:"pageSize" validate:"required,number,lt=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDAtPathReq
type IDAtPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base ID response data | 基础ID信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id *uint64 `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础UUID信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id *string `json:"id,optional"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The response data of spu information | Spu信息
// swagger:model SpuInfo
type SpuInfo struct {
	BaseIDInfo
	// Status
	Status *uint32 `json:"status,optional"`
	// Sort
	Sort *uint32 `json:"sort,optional"`
	// Name
	Name *string `json:"name,optional"`
	// Keyword
	Keyword *string `json:"keyword,optional"`
	// Introduction
	Introduction *string `json:"introduction,optional"`
	// Description
	Description *string `json:"description,optional"`
	// BarCode
	BarCode *string `json:"barCode,optional"`
	// CategoryId
	CategoryId *uint64 `json:"categoryId,optional"`
	// BrandId
	BrandId *uint64 `json:"brandId,optional"`
	// PicUrl
	PicUrl *string `json:"picUrl,optional"`
	// SliderPicUrls
	SliderPicUrls *string `json:"sliderPicUrls,optional"`
	// VideoUrl
	VideoUrl *string `json:"videoUrl,optional"`
	// Unit
	Unit *uint32 `json:"unit,optional"`
	// SpecType
	SpecType *bool `json:"specType,optional"`
	// Price
	Price *int32 `json:"price,optional"`
	// MarketPrice
	MarketPrice *int32 `json:"marketPrice,optional"`
	// CostPrice
	CostPrice *int32 `json:"costPrice,optional"`
	// Stock
	Stock *int32 `json:"stock,optional"`
	// SalesCount
	SalesCount *int32 `json:"salesCount,optional"`
	// VirtualSalesCount
	VirtualSalesCount *int32 `json:"virtualSalesCount,optional"`
	// BrowseCount
	BrowseCount *int32 `json:"browseCount,optional"`
}

// swagger:model SkuInfo
type SkuInfo struct {
	// Price
	Price *int32 `json:"price,optional"`
	// MarketPrice
	MarketPrice *int32 `json:"marketPrice,optional"`
	// CostPrice
	CostPrice *int32 `json:"costPrice,optional"`
	// BarCode
	BarCode *string `json:"barCode,optional"`
	// PicUrl
	PicUrl *string `json:"picUrl,optional"`
	// Stock
	Stock *int32 `json:"stock,optional"`
	// Weight
	Weight *float64 `json:"weight,optional"`
	// Volume
	Volume *float64 `json:"volume,optional"`
	// Properties
	Properties []*SkuProperty `json:"properties,optional"`
}

type SpuBase struct {
	// Name 商品名称
	Name string `json:"name"`
	// Keyword 关键字
	Keyword string `json:"keyword"`
	// Introduction 商品简介
	Introduction string `json:"introduction"`
	// Description 商品详情
	Description string `json:"description"`
	// CategoryId 商品分类编号
	CategoryId uint64 `json:"categoryId"`
	// BrandId 商品品牌编号
	BrandId uint64 `json:"brandId"`
	// PicUrl 商品封面图
	PicUrl string `json:"picUrl"`
	// SliderPicUrls 商品轮播图
	SliderPicUrls []string `json:"sliderPicUrls,optional"`
	// VideoUrl 商品视频
	VideoUrl *string `json:"videoUrl,optional"`
	// Unit 单位
	Unit uint32 `json:"unit"`
	// Sort 排序字段
	Sort uint32 `json:"sort"`
	// SpecType 规格类型
	SpecType bool `json:"specType"`
	// VirtualSalesCount 虚拟销量
	VirtualSalesCount *int32 `json:"virtualSalesCount,optional"`
}

type SkuBase struct {
	// Price 销售价格，单位：分
	Price int32 `json:"price"`
	// MarketPrice 市场价
	MarketPrice *int32 `json:"marketPrice,optional"`
	// CostPrice 成本价
	CostPrice *int32 `json:"costPrice,optional"`
	// BarCode 条形码
	BarCode *string `json:"barCode,optional"`
	// PicUrl 图片地址
	PicUrl string `json:"picUrl"`
	// Stock 库存
	Stock int32 `json:"stock"`
	// Weight 商品重量,单位：kg 千克
	Weight *float64 `json:"weight,optional"`
	// Volume 商品体积,单位：m^3 平米
	Volume *float64 `json:"volume,optional"`
	// Properties 属性数组
	Properties []SkuProperty `json:"properties,optional"`
}

// swagger:model SpuCreateReq
type SpuCreateReq struct {
	SpuBase
	// skus
	Skus []SkuBase `json:"skus"`
}

// swagger:model SpuDetailResp
type SpuDetailResp struct {
	IDReq
	SpuInfo
	//		// SalesCount
	//		SalesCount int32 `json:"salesCount"`
	//		// BrowseCount
	//		BrowseCount int32 `json:"browseCount"`
	//		// Status
	//		Status *uint32 `json:"status,optional"`
	// skus
	Skus []SkuInfo `json:"skus"`
}

// swagger:model SpuPageReq
type SpuPageReq struct {
	PageInfo
	// Name 商品名称
	Name *string `json:"name,optional"`
	// TabType 前端请求的tab类型
	TabType *int32 `json:"tabType,optional"`
	// CategoryId 商品分类编号
	CategoryId *uint64 `json:"categoryId,optional"`
	// CreateTime 创建时间
	CreateTime []*uint64 `json:"createTime,optional"`
}

// swagger:model SpuResp
type SpuResp struct {
	SpuInfo
	IDReq
	// Price 商品价格
	Price *int32 `json:"price,optional"`
	// MarketPrice 市场价，单位使用：分
	MarketPrice *int32 `json:"marketPrice,optional"`
	// CostPrice 成本价，单位使用：分
	CostPrice *int32 `json:"costPrice,optional"`
	// Stock 商品库存
	Stock *int32 `json:"stock,optional"`
	// CreateTime 商品创建时间
	CreateTime *uint64 `json:"createTime,optional"`
	// Status 商品状态
	Status *uint32 `json:"status,optional"`
}

// swagger:model SpuSimpleResp
type SpuSimpleResp struct {
	IDReq
	// Name 商品名称
	Name *string `json:"name,optional"`
	// Price 商品价格，单位使用：分
	Price *int32 `json:"price,optional"`
	// MarketPrice 商品市场价，单位使用：分
	MarketPrice *int32 `json:"marketPrice,optional"`
	// CostPrice 商品成本价，单位使用：分
	CostPrice *int32 `json:"costPrice,optional"`
	// Stock 商品库存
	Stock *int32 `json:"stock,optional"`
	// SalesCount 商品销量
	SalesCount *int32 `json:"salesCount,optional"`
	// VirtualSalesCount 商品虚拟销量
	VirtualSalesCount *int32 `json:"virtualSalesCount,optional"`
	// BrowseCount 商品浏览量
	BrowseCount *int32 `json:"browseCount,optional"`
}

// swagger:model SpuSimpleListResp
type SpuSimpleListResp struct {
	BaseListInfo
	Data []SpuSimpleResp `json:"data"`
}

// swagger:model SpuListResp
type SpuListResp struct {
	BaseListInfo
	Data []SpuDetailResp `json:"data"`
}

// swagger:model SpuCountResp
type SpuCountResp struct {
	Data map[int32]uint64 `json:"data"`
}

// swagger:model SpuPageResp
type SpuPageResp struct {
	BaseListInfo
	Data []SpuResp `json:"data"`
}

// swagger:model SpuUpdateReq
type SpuUpdateReq struct {
	IDReq
	SpuBase
	// SalesCount 商品销量
	SalesCount *int32 `json:"salesCount,optional"`
	// BrowseCount 浏览量
	BrowseCount *int32 `json:"browseCount,optional"`
	// Status 商品状态
	Status *uint32 `json:"status,optional"`
	// skus SKU 数组
	Skus []SkuBase `json:"skus"`
}

// swagger:model SpuUpdateStatusReq
type SpuUpdateStatusReq struct {
	IDReq
	// Status
	Status uint32 `json:"status"`
}

type SkuProperty struct {
	// PropertyId
	PropertyId uint64 `json:"propertyId"`
	// PropertyName
	PropertyName string `json:"propertyName"`
	// ValueId
	ValueId uint64 `json:"valueId"`
	// ValueName
	ValueName string `json:"valueName"`
}

type BrandBase struct {
	// Name 品牌名称
	Name string `json:"name"`
	// PicUrl 品牌图片
	PicUrl string `json:"picUrl"`
	// Sort 品牌排序
	Sort uint32 `json:"sort"`
	// Description 品牌描述
	Description *string `json:"description,optional"`
	// Status 状态
	Status uint32 `json:"status"`
}

// swagger:model BrandCreateReq
type BrandCreateReq struct {
	BrandBase
}

// swagger:model BrandListReq
type BrandListReq struct {
	// Name 品牌名称
	Name *string `path:"name,optional"`
}

// swagger:model BrandPageReq
type BrandPageReq struct {
	PageAtPathInfo
	// Name 品牌名称
	Name *string `path:"name,optional"`
	// Status 状态
	Status *uint32 `path:"status,optional"`
	// Date 创建时间
	CreatedAt []int64 `path:"createdAt,optional"`
}

// swagger:model BrandListInfoResp
type BrandListInfoResp struct {
	BaseListInfo
	Data []BrandResp `json:"data"`
}

// swagger:model BrandResp
type BrandResp struct {
	BaseIDInfo
	// 品牌id
	Id *uint64 `json:"id"`
	// Name 品牌名称
	Name *string `json:"name"`
	// PicUrl 品牌图片
	PicUrl *string `json:"picUrl"`
	// Sort 品牌排序
	Sort *uint32 `json:"sort"`
	// Description 品牌描述
	Description *string `json:"description"`
	// Status 状态
	Status *uint32 `json:"status"`
}

// swagger:model BrandUpdateReq
type BrandUpdateReq struct {
	BaseIDInfo
	BrandBase
}

// The response data of category information | Category信息
// swagger:model CategoryInfo
type CategoryInfo struct {
	BaseIDInfo
	// Status 开启状态
	Status *uint32 `json:"status,optional"`
	// ParentId 父分类编号
	ParentId *uint64 `json:"parentId,optional"`
	// Name 分类名称
	Name *string `json:"name,optional"`
	// PicUrl 移动端分类图
	PicUrl *string `json:"picUrl,optional"`
	// BigPicUrl PC 端分类图
	BigPicUrl *string `json:"bigPicUrl,optional"`
}

type CategoryBase struct {
	// ParentId 父分类编号
	ParentId uint64 `json:"parentId"`
	// Name 分类名称
	Name string `json:"name"`
	// PicUrl 移动端分类图
	PicUrl string `json:"picUrl"`
	// BigPicUrl PC 端分类图
	BigPicUrl *string `json:"bigPicUrl,optional"`
	// Sort 分类排序
	Sort *uint32 `json:"sort,optional"`
	// Status 开启状态
	Status uint32 `json:"status"`
}

// swagger:model CategoryCreateReq
type CategoryCreateReq struct {
	CategoryBase
	Description *string `json:"description,optional"`
}

// swagger:model CategoryListReq
type CategoryListReq struct {
	// ParentId 父分类编号
	ParentId *uint64 `path:"parentId,optional"`
	// Name 分类名称
	Name *string `path:"name,optional"`
	// Status 开启状态
	Status *uint32 `path:"status,optional"`
}

// swagger:model CategoryResp
type CategoryResp struct {
	BaseIDInfo
	// Status 开启状态
	Status *uint32 `json:"status,optional"`
	// ParentId 父分类编号
	ParentId *uint64 `json:"parentId,optional"`
	// Name 分类名称
	Name *string `json:"name,optional"`
	// PicUrl 移动端分类图
	PicUrl *string `json:"picUrl,optional"`
	// BigPicUrl PC 端分类图
	BigPicUrl *string `json:"bigPicUrl,optional"`
	// Sort 分类排序
	Sort *uint32 `json:"sort,optional"`
}

// swagger:model CategoryUpdateReq
type CategoryUpdateReq struct {
	IDReq
	CategoryBase
	Description string `json:"description,optional"`
}

// swagger:model CategoryListResp
type CategoryListResp struct {
	BaseListInfo
	Data []CategoryResp `json:"data"`
}

// The response data of comment information | Comment信息
// swagger:model CommentInfo
type CommentInfo struct {
	BaseIDInfo
	// UserId
	UserId *int64 `json:"userId,optional"`
	// UserNickname
	UserNickname *string `json:"userNickname,optional"`
	// UserAvatar
	UserAvatar *string `json:"userAvatar,optional"`
	// Anonymous
	Anonymous *bool `json:"anonymous,optional"`
	// OrderId
	OrderId *int64 `json:"orderId,optional"`
	// OrderItemId
	OrderItemId *int64 `json:"orderItemId,optional"`
	// SpuId
	SpuId *int64 `json:"spuId,optional"`
	// SpuName
	SpuName *string `json:"spuName,optional"`
	// SkuId
	SkuId *int64 `json:"skuId,optional"`
	// SkuPicUrl
	SkuPicUrl *string `json:"skuPicUrl,optional"`
	// SkuProperties
	SkuProperties *string `json:"skuProperties,optional"`
	// Visible
	Visible *bool `json:"visible,optional"`
	// Scores
	Scores *int32 `json:"scores,optional"`
	// DescriptionScores
	DescriptionScores *int32 `json:"descriptionScores,optional"`
	// BenefitScores
	BenefitScores *int32 `json:"benefitScores,optional"`
	// Content
	Content *string `json:"content,optional"`
	// PicUrls
	PicUrls *string `json:"picUrls,optional"`
	// ReplyStatus
	ReplyStatus *bool `json:"replyStatus,optional"`
	// ReplyUserId
	ReplyUserId *int64 `json:"replyUserId,optional"`
	// ReplyContent
	ReplyContent *string `json:"replyContent,optional"`
	// ReplyTime
	ReplyTime *int64 `json:"replyTime,optional"`
}

// The response data of comment list | Comment列表数据
// swagger:model CommentListResp
type CommentListResp struct {
	BaseDataInfo
	// Comment list data | Comment列表数据
	Data CommentListInfo `json:"data"`
}

// Comment list data | Comment列表数据
// swagger:model CommentListInfo
type CommentListInfo struct {
	BaseListInfo
	// The API list data | Comment列表数据
	Data []CommentInfo `json:"data"`
}

// Get comment list request params | Comment列表请求参数
// swagger:model CommentListReq
type CommentListReq struct {
	PageInfo
	// UserNickname
	UserNickname *string `json:"userNickname,optional"`
	// UserAvatar
	UserAvatar *string `json:"userAvatar,optional"`
	// SpuName
	SpuName *string `json:"spuName,optional"`
}

// Comment information response | Comment信息返回体
// swagger:model CommentInfoResp
type CommentInfoResp struct {
	BaseDataInfo
	// Comment information | Comment数据
	Data CommentInfo `json:"data"`
}

// The response data of property information | Property信息
// swagger:model PropertyInfo
type PropertyInfo struct {
	BaseIDInfo
	// Status 状态
	Status *uint32 `json:"status,optional"`
	// Name 名称
	Name *string `json:"name,optional"`
	// Remark 备注
	Remark *string `json:"remark,optional"`
}

type PropertyBase struct {
	// Name 名称
	Name string `json:"name"`
	// Remark 备注
	Remark *string `json:"remark,optional"`
}

// swagger:model PropertyListReq
type PropertyListReq struct {
	// Name 名称
	Name *string `json:"name,optional"`
}

// swagger:model PropertyPageReq
type PropertyPageReq struct {
	PageInfo
	// Status 状态
	Status *uint32 `json:"status,optional"`
	// Name 名称
	Name *string `json:"name,optional"`
	// CreatedAt 创建时间
	CreatedAt []int64 `json:"createdAt,optional"`
}

// swagger:model PropertyResp
type PropertyResp struct {
	PropertyInfo
	IDReq
}

// swagger:model PropertyListResp
type PropertyListResp struct {
	BaseListInfo
	Data []PropertyInfo `json:"data"`
}

// swagger:model PropertyUpdateReq
type PropertyUpdateReq struct {
	PropertyBase
	IDReq
}

// swagger:model PropertyAndValueListResp
type PropertyAndValueListResp struct {
	PropertyAndValue []*PropertyAndValue `json:"propertyAndValue"`
}

type PropertyAndValue struct {
	IDReq
	// Name 名称
	Name *string `json:"name"`
	// 属性值的集合
	Value []*Value `json:"value"`
}

type Value struct {
	Id *uint64 `json:"id"`
	// Name
	Name *string `json:"name"`
}

// The response data of property value information | PropertyValue信息
// swagger:model PropertyValueInfo
type PropertyValueInfo struct {
	BaseIDInfo
	// Status 状态
	Status *uint32 `json:"status,optional"`
	// PropertyId 属性项的编号
	PropertyId *uint64 `json:"propertyId,optional"`
	// Name 名称
	Name *string `json:"name,optional"`
	// Remark 备注
	Remark *string `json:"remark,optional"`
}

type PropertyValueBase struct {
	// PropertyId 属性项的编号
	PropertyId uint64 `json:"propertyId"`
	// Name 名称
	Name string `json:"name"`
	// Remark 备注
	Remark *string `json:"remark,optional"`
}

// swagger:model PropertyValueDetailResp
type PropertyValueDetailResp struct {
	// PropertyId 属性的编号
	PropertyId uint64 `json:"propertyId"`
	// Name 属性的名称
	PropertyName string `json:"propertyName"`
	// Remark 属性值的编号
	ValueId uint64 `json:"valueId"`
	// Name 属性值的名称
	ValueName string `json:"valueName"`
}

// swagger:model PropertyValuePageReq
type PropertyValuePageReq struct {
	PageInfo
	// PropertyId 属性的编号
	PropertyId *uint64 `json:"propertyId,optional"`
	// Name 名称
	Name *string `json:"name,optional"`
	// Status 状态
	Status *uint32 `json:"status,optional"`
}

// swagger:model PropertyValueResp
type PropertyValueResp struct {
	PropertyValueBase
	IDReq
	// 创建时间
	CreateTime []*uint64 `json:"createTime"`
}

// swagger:model PropertyValueListResp
type PropertyValueListResp struct {
	BaseListInfo
	Data []PropertyValueInfo `json:"data"`
}

// swagger:model PropertyValueUpdateReq
type PropertyValueUpdateReq struct {
	PropertyValueBase
	IDReq
}
