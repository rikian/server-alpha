package trash

// func registerService(s *grpc.Server) {
// 	pb.RegisterUserRPCServer(s, &users.UserImpl{})
// pb.RegisterProductRPCServer(s, &ProductService{})
// pb.RegisterImageRPCServer(s, &ImageService{})
// }

// func run(address string) error {
// 	listen, err := net.Listen("tcp", address)

// 	if err != nil {
// 		log.Fatal("failed listen grpc server")
// 	}

// 	server := grpc.NewServer()
// 	registerService(server)
// 	return server.Serve(listen)
// }

// func ListenAndServe(address string) {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error when loading .env file")
// 	}

// 	config.ConnectDB()

// 	log.Print("Grpc service 1 running at " + address)

// 	if err := run(address); err != nil {
// 		log.Fatalf("Failed running grpc server at %v", address)
// 	}
// }

// var UserRepo = repository.InitUserRepo()
// var authService = users.NewUserService()
// var productService products.ProductService = products.NewProductService()

// type UserService struct {
// 	pb.UnimplementedUserRPCServer
// }

// type ProductService struct {
// 	pb.UnimplementedProductRPCServer
// }

// type ImageService struct {
// 	pb.UnimplementedImageRPCServer
// }

// // images handler
// func (s *ImageService) InsertProductId(ctx context.Context, input *pb.DataProduct) (*pb.ResponseInsertProductID, error) {
// 	return UserRepo.InsertProductId(ctx, input)
// }

// // auth handler
// func (s *UserService) LoginUser(ctx context.Context, input *pb.DataLogin) (*pb.ResponseLogin, error) {
// 	return authService.ServiceLogin(ctx, input)
// }

// func (s *UserService) RegisterUser(ctx context.Context, input *pb.DataRegister) (*pb.ResponseRegister, error) {
// 	return authService.ServiceRegister(ctx, input)
// }

// func (s *UserService) SelectSessionUserById(ctx context.Context, input *pb.DataSession) (*pb.ResponseSession, error) {
// 	return authService.GetSession(ctx, input)
// }

// func (s *UserService) SelectUser(ctx context.Context, input *pb.DataSelectUser) (*pb.ResponseSelectUser, error) {
// 	return UserRepo.SelectUser(ctx, input)
// }

// func (s *UserService) InsertProduct(ctx context.Context, input *pb.DataInsertProduct) (*pb.ResponseInsertProduct, error) {
// 	return UserRepo.InsertProduct(ctx, input)
// }

// func (s *UserService) DeleteProduct(ctx context.Context, input *pb.DataDeleteProduct) (*pb.ResponseDeleteProduct, error) {
// 	return UserRepo.DeleteProduct(ctx, input)
// }

// func (s *UserService) UpdateProduct(ctx context.Context, input *pb.DataUpdateProduct) (*pb.ResponseUpdateProduct, error) {
// 	return UserRepo.UpdateProduct(ctx, input)
// }

// func (s *ProductService) ProductsRPC(ctx context.Context, input *pb.User) (*pb.Products, error) {
// 	return productService.ProductsRPC(ctx, input)
// }

// func (s *ProductService) SingelProduct(ctx context.Context, input *pb.RequestProduct) (*pb.Product, error) {
// 	return productService.SingelProduct(ctx, input)
// }
