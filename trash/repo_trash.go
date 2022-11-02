package trash

// func (u *userRepo) InsertProductId(ctx context.Context, input *pb.DataProduct) (*pb.ResponseInsertProductID, error) {
// 	var db gorm.DB = *config.DB

// 	u.DataInsertProduct = &entities.Product{
// 		UserId:       input.UserId,
// 		ProductId:    input.ProductId,
// 		ProductName:  "",
// 		ProductImage: input.ProductImage,
// 		ProductInfo:  "",
// 		ProductStock: 0,
// 		ProductPrice: 0,
// 		ProductSell:  0,
// 		CreatedDate:  input.CreatedDate,
// 		LastUpdate:   "",
// 	}

// 	// begin
// 	tx := db.Begin()

// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// 	// logic
// 	insertProduct := tx.Create(&u.DataInsertProduct)

// 	if insertProduct.Error != nil {
// 		log.Print(insertProduct.Error.Error())
// 		tx.Rollback()
// 		return nil, insertProduct.Error
// 	}

// 	// comit
// 	comit := tx.Commit()

// 	if comit.Error != nil {
// 		log.Print(comit.Error.Error())
// 		return nil, comit.Error
// 	}

// 	u.ResponseInsertProductId = &pb.ResponseInsertProductID{
// 		Status: "ok",
// 		Error:  "null",
// 	}

// 	return u.ResponseInsertProductId, nil
// }

// func (u *userRepo) InsertProduct(ctx context.Context, i *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error) {
// 	log.Print(i)
// 	var db gorm.DB = *config.DB
// 	var time string = time.Now().Format("20060102150405")
// 	// begin
// 	tx := db.Begin()
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// 	// logic
// 	product := &entities.Product{
// 		UserId:       i.UserId,
// 		ProductId:    i.ProductId,
// 		ProductName:  i.ProductName,
// 		ProductInfo:  i.ProductInfo,
// 		ProductStock: int(i.ProductStock),
// 		ProductPrice: int(i.ProductPrice),
// 		ProductSell:  int(i.ProductSell),
// 		LastUpdate:   time,
// 	}

// 	updateProduct := tx.Model(product).Updates(product).Where("user_id = ? AND product_id=?", i.UserId, i.ProductId)

// 	if updateProduct.Error != nil {
// 		log.Print(updateProduct.Error)
// 		tx.Rollback()
// 		return nil, updateProduct.Error
// 	}

// 	err := tx.Model(u.product).Where("product_id = ?", i.ProductId).First(product).Error

// 	if err != nil {
// 		log.Print(err)
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	// comit
// 	comit := tx.Commit()

// 	if comit.Error != nil {
// 		log.Print(comit.Error)
// 		return nil, comit.Error
// 	}

// 	i.LastUpdate = product.LastUpdate
// 	i.ProductImage = product.ProductImage

// 	return &pb.ResponseInsertProduct{
// 		Status:  200,
// 		Message: "ok",
// 		Product: i,
// 	}, nil
// }

// func (u *userRepo) DeleteProduct(ctx context.Context, i *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error) {
// 	var db gorm.DB = *config.DB

// 	// begin
// 	tx := db.Begin()

// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// 	// logic
// 	deleteProduct := tx.Where("user_id = ? AND product_id=?", i.UserId, i.ProductId).Delete(&u.product)

// 	if deleteProduct.Error != nil {
// 		log.Print(deleteProduct.Error.Error())
// 		tx.Rollback()
// 		return nil, deleteProduct.Error
// 	}

// 	// delete image product from directory server image
// 	deleteDirImage, err := http.DeleteImage(i.UserId, i.ProductId)

// 	if err != nil {
// 		log.Print(err.Error())
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	if deleteDirImage.Message != "ok" {
// 		log.Print(deleteDirImage.Message)
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	// comit
// 	comit := tx.Commit()

// 	if comit.Error != nil {
// 		log.Print(comit.Error.Error())
// 		return nil, comit.Error
// 	}

// 	return &pb.ResponseDeleteProduct{
// 		Status:    200,
// 		Message:   "ok",
// 		ProductId: i.ProductId,
// 	}, nil
// }

// func (u *userRepo) UpdateProduct(ctx context.Context, i *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
// 	var db gorm.DB = *config.DB
// 	var time string = time.Now().Format("20060102150405")
// 	// begin
// 	tx := db.Begin()
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// 	// logic
// 	product := &entities.Product{
// 		UserId:       i.UserId,
// 		ProductId:    i.ProductId,
// 		ProductName:  i.ProductName,
// 		ProductImage: i.ProductImage,
// 		ProductInfo:  i.ProductInfo,
// 		ProductStock: int(i.ProductStock),
// 		ProductPrice: int(i.ProductPrice),
// 		ProductSell:  int(i.ProductSell),
// 		LastUpdate:   time,
// 	}

// 	updateProduct := tx.Model(product).Updates(product).Where("user_id = ? AND product_id=?", i.UserId, i.ProductId)

// 	if updateProduct.Error != nil {
// 		log.Print(updateProduct.Error)
// 		tx.Rollback()
// 		return nil, updateProduct.Error
// 	}

// 	// comit
// 	comit := tx.Commit()

// 	if comit.Error != nil {
// 		log.Print(comit.Error)
// 		return nil, comit.Error
// 	}

// 	i.LastUpdate = product.LastUpdate

// 	return &pb.ResponseUpdateProduct{
// 		Status:  200,
// 		Message: "ok",
// 		Product: i,
// 	}, nil
// }

// func UpdateSession(userId, session string) {}

// func (u *userRepo) SelectUsers(ctx context.Context, input *pb.Limit) (*pb.DataUsers, error) {
// 	var db gorm.DB = *config.DB

// 	u.User = &entities.User{}
// 	u.Users = []entities.User{}

// 	err := db.Model(u.User).Preload("UserStatus").Preload("Products").Find(&u.Users)

// 	if err.Error != nil {
// 		log.Print(err.Error.Error())
// 		return nil, errors.New("something wrong")
// 	}

// 	u.GrpcUsers = &pb.DataUsers{}

// 	for i := 0; i < len(u.Users); i++ {
// 		u.GrpcUser = &pb.DataUser{
// 			UserId:      u.Users[i].UserId,
// 			UserEmail:   u.Users[i].UserEmail,
// 			UserName:    u.Users[i].UserName,
// 			UserImage:   u.Users[i].UserImage,
// 			UserStatus:  u.Users[i].UserStatus.Status,
// 			CreatedDate: u.Users[i].CreatedDate,
// 			LastUpdate:  u.Users[i].LastUpdate,
// 		}

// 		for j := 0; j < len(u.Users[i].Products); j++ {
// 			u.GrpcUser.Products = append(u.GrpcUser.Products, &pb.Product{
// 				UserId:       u.Users[i].Products[j].UserId,
// 				ProductId:    u.Users[i].Products[j].ProductId,
// 				ProductName:  u.Users[i].Products[j].ProductName,
// 				ProductPrice: uint32(u.Users[i].Products[j].ProductPrice),
// 				ProductStock: uint32(u.Users[i].Products[j].ProductStock),
// 				CreatedDate:  u.Users[i].Products[j].CreatedDate,
// 				LastUpdate:   u.Users[i].Products[j].LastUpdate,
// 			})
// 		}

// 		u.GrpcUsers.Data = append(u.GrpcUsers.Data, u.GrpcUser)
// 	}

// 	return u.GrpcUsers, nil
// }

// var db gorm.DB = *config.DB
// 	var time string = time.Now().Format("20060102150405")

// 	u.DataInsertProduct = &entities.Product{
// 		UserId:       i.UserId,
// 		ProductId:    uuid.New().String(),
// 		ProductName:  i.ProductName,
// 		ProductImage: i.ProductImage,
// 		ProductInfo:  i.ProductInfo,
// 		ProductStock: int(i.ProductStock),
// 		ProductPrice: int(i.ProductPrice),
// 		ProductSell:  int(i.ProductSell),
// 		CreatedDate:  time,
// 		LastUpdate:   time,
// 	}

// 	// begin
// 	tx := db.Begin()

// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	if tx.Error != nil {
// 		log.Print(tx.Error.Error())
// 		return nil, tx.Error
// 	}

// logic
// insertProduct := tx.Create(&u.DataInsertProduct)

// if insertProduct.Error != nil {
// 	log.Print(insertProduct.Error.Error())
// 	tx.Rollback()
// 	return nil, insertProduct.Error
// }

// save image

// comit
// comit := tx.Commit()

// if comit.Error != nil {
// 	log.Print(comit.Error.Error())
// 	return nil, comit.Error
// }

// i.ProductId = u.DataInsertProduct.ProductId
// i.CreatedDate = u.DataInsertProduct.CreatedDate
// i.LastUpdate = u.DataInsertProduct.LastUpdate

// return &pb.ResponseInsertProduct{
// 	Status:  200,
// 	Message: "ok",
// 	Product: i,
// }, nil
