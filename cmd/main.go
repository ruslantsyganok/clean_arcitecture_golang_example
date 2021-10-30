package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"zen_api/internal/app"
	"zen_api/internal/service"

	"zen_api/internal/repository"
	desc "zen_api/pkg"
)

func main() {
	// DB
	db, err := repository.NewDB()
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	// preparing config file
	viper.AddConfigPath("../config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read from a config")
	}

	// yandex cloud storage
	host := viper.Get("storage.host").(string)
	bucketName := viper.Get("storage.bucketName").(string)
	region := viper.Get("storage.region").(string)
	keyID := viper.Get("storage.keyID").(string)
	keySecret := viper.Get("storage.keySecret").(string)

	// JWT
	signedKeyJWT := viper.Get("jwt.signedKey").(string)
	tokenManager := service.NewTokenManager(signedKeyJWT)

	// Register all services
	dao := repository.NewDAO(db)
	userService := service.NewUserService(dao)
	answerService := service.NewAnswerService(dao)
	authService := service.NewAuthService(dao, tokenManager)
	courseService := service.NewCourseService(dao)
	questionService := service.NewQuestionService(dao)
	reviewService := service.NewReviewService(dao)
	scoreService := service.NewScoreService(dao)
	sectionService := service.NewSectionService(dao)
	indicatorService := service.NewIndicatorService(dao)
	pollService := service.NewPollService(dao)
	emailVerificationService := service.NewEmailVerificationService(dao)
	fileUploaderService := service.NewFileUploaderService(dao, bucketName, host, region, keyID, keySecret)

	// Payment service (qiwi)
	paymentSecret := viper.Get("payment.secret").(string)
	paymentService := service.NewPaymentService(paymentSecret)

	// Interceptors
	grpcOpts := app.GrpcInterceptor()
	httpOpts := app.HttpInterceptor()

	// Starting gRPC server
	go func() {
		listener, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			log.Fatalln(err)
		}

		grpcServer := grpc.NewServer(grpcOpts)
		desc.RegisterMicroserviceServer(grpcServer, app.NewMicroservice(
			userService,
			answerService,
			authService,
			courseService,
			questionService,
			reviewService,
			scoreService,
			sectionService,
			indicatorService,
			pollService,
			emailVerificationService,
			fileUploaderService,
			tokenManager,
			paymentService))

		err = grpcServer.Serve(listener)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Starting HTTP server
	mux := runtime.NewServeMux(httpOpts)
	err = desc.RegisterMicroserviceHandlerServer(context.Background(), mux, app.NewMicroservice(
		userService,
		answerService,
		authService,
		courseService,
		questionService,
		reviewService,
		scoreService,
		sectionService,
		indicatorService,
		pollService,
		emailVerificationService,
		fileUploaderService,
		tokenManager,
		paymentService))
	if err != nil {
		log.Println("cannot register this service")
	}
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
