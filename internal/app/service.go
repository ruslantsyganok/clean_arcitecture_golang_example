package app

import (
	"zen_api/internal/service"
	desc "zen_api/pkg"
)

type MicroserviceServer struct {
	desc.UnimplementedMicroserviceServer
	userService         service.UserService
	answerService       service.AnswerService
	authService         service.AuthService
	courseService       service.CourseService
	questionService     service.QuestionService
	reviewService       service.ReviewService
	scoreService        service.ScoreService
	sectionService      service.SectionService
	indicatorService    service.IndicatorService
	pollService         service.PollService
	emailService        service.EmailService
	fileUploaderService service.FileUploaderService
	tokenManager        service.TokenManager
	paymentService      service.PaymentService
}

func NewMicroservice(
	userService service.UserService,
	answerService service.AnswerService,
	authService service.AuthService,
	courseService service.CourseService,
	questionService service.QuestionService,
	reviewService service.ReviewService,
	scoreService service.ScoreService,
	sectionService service.SectionService,
	indicatorService service.IndicatorService,
	pollService service.PollService,
	emailService service.EmailService,
	fileUploaderService service.FileUploaderService,
	tokenManager service.TokenManager,
	paymentService service.PaymentService) *MicroserviceServer {
	return &MicroserviceServer{
		userService:         userService,
		answerService:       answerService,
		authService:         authService,
		courseService:       courseService,
		questionService:     questionService,
		reviewService:       reviewService,
		scoreService:        scoreService,
		sectionService:      sectionService,
		indicatorService:    indicatorService,
		pollService:         pollService,
		emailService:        emailService,
		fileUploaderService: fileUploaderService,
		tokenManager:        tokenManager,
		paymentService:      paymentService,
	}
}
