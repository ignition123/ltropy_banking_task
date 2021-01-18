package Web

import(
	"cns"
	"controllers/auth"
	"controllers/brokers"
	"lib"
	"middlewares"
)

func Routes(){

	defer lib.HandlePanic()

	httpApp := cns.Http{}

	// auth services

	/////////////////////////////////////////////////////////////// admin auth apis ///////////////////////////////////////////////////////////////

	// Admin Registration
	httpApp.Post("/admin/registration", auth.Register).Middlewares(middlewares.ValidateSession)

	// Admin SignIn
	httpApp.Post("/admin/signin", auth.SignIn)

	// password reset
	httpApp.Post("/admin/password_reset", auth.ResetPassword).Middlewares(middlewares.ValidateSession)

	// password update route
	httpApp.Post("/admin/password_update", auth.UpdatePassword).Middlewares(middlewares.ValidateSession)

	// update permission
	httpApp.Post("/admin/update_permission", auth.UpdatePermission).Middlewares(middlewares.ValidateSession)

	// update profile
	httpApp.Post("/admin/update_profile", auth.UpdateProfile).Middlewares(middlewares.ValidateSession)

	// create superAdmin
	httpApp.Post("/admin/add_super_admin", auth.CreateSuperAdmin).Middlewares(middlewares.ValidateSession)

	// Get profile
	httpApp.Get("/admin/get_profile", auth.GetProfile).Middlewares(middlewares.ValidateSession)

	////////////////////////////////////////////////////////// brokers apis ///////////////////////////////////////////////////////////////////

	// approve broker contacts
	httpApp.Post("/broker/approve_contacts", brokers.ApproveContacts).Middlewares(middlewares.ValidateSession)

	// freeze broker
	httpApp.Post("/broker/freeze_broker_account", brokers.FreezeAccount).Middlewares(middlewares.ValidateSession)

	// create new interactive broker id
	httpApp.Post("/broker/create_new_interactive_id", brokers.CreateNewInteractiveId).Middlewares(middlewares.ValidateSession)

	// create new broadcast broker id
	httpApp.Post("/broker/create_new_broadcast_id", brokers.CreateNewBroadCastId).Middlewares(middlewares.ValidateSession)

	// trading brokers id lists
	httpApp.Post("/broker/get_broker_ids", brokers.GetAllBrokerIds).Middlewares(middlewares.ValidateSession)
}