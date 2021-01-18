package messages

const(
	// General Messages
	ErrorMessage = "Oops, something went wrong please try after sometimes..."
	RequestValidateFail = "Invalid request..."

	EmailEmpty = "Email address cannot be empty..."
	NameEmpty = "Name cannot be empty..."
	SuperAdminFlagEmpty = "Super admin flag cannot be empty..."
	CreateAdminFlagEmpty = "Create admin flag cannot be empty..."
	KYCApproveContactEmpty = "KYC approve contact cannot be empty..."
	KYCApproveEmpty = "KYC complete permissions cannot be empty..."
	AccountBlockedEmpty = "Account blocked permissions cannot be empty..."
	UpdateProfileEmpty = "Update profile cannot be empty..."
	FreezeAccountEmpty = "Freeze account permission cannot be empty..."
	CreateNewBrokerIdEmpty = "Broker id creation permission cannot be empty..."
	ResettingPasswordEmpty = "Broker password reset permission cannot be empty..."
	ExchangePushEmpty = "Exchange push empty cannot be empty..."

	UserAlreadyExists = "User already exists..."
	RegisterSuccess = "Account created successfully..."
	PasswordEmpty = "Password cannot be empty..."

	InvalidUserIdPassword = "Invalid email or password..."
	DontHavePermission = "You don't have permission for the following request..."
	PasswordResetSuccess = "Password reset successfully..."
	PasswordAlreadyDefault = "Password is already set to default..."
	ConfirmPasswordEmpty = "Confirm password cannot be empty..."
	PasswordDoesNotMatch = "Password does not match with confirm password..."
	PasswordUpdatedSuccess = "Password updated successfully..."
	PermissionUpdated = "Permission updated successfully..."
	UpdateProfileSuccess = "Profile updated successfully..."
	InvalidDOB = "Invalid date of birth..."
	CreatedSuperAdminSuccess = "Super admin created successfully..."

	CompanyIdEmpty        = "Company id cannot be empty..."
	InvalidPassword       = "Invalid password..."
	BrokerContactApprovedFailed = "Failed to approve contact of the broker..."
	BrokerContactApproved = "Broker contact approved..."
	BrokerAccountFroze    = "Broker account froze successfully..."
	FailedToFrozeAccount = "Failed to freeze account of the broker..."
	BrokerNewInteractiveIdCreated = "Broker new interactive id created successfully..."
	BrokerNewInteractiveIdFailed = "Failed to create interactive id..."
	BrokerNewBroadcastIdFailed = "Failed to create broadcast id..."
	BrokerNewBroadcastIdCreated = "Broker new broadcast id created successfully..."

	NoTradingIds = "No trading ids created..."
)