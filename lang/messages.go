package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var translationMessages = []i18n.Message{
	{
		ID:    "site_name",
		Other: "Golang Base Project",
	},
	{
		ID:    "home",
		Other: "Home",
	},
	{
		ID:    "activation_validation_token",
		Other: "Please provide a valid activation token",
	},
	{
		ID:    "activation_success",
		Other: "Account activated. You may now proceed to login to your account.",
	},
	{
		ID:    "activate",
		Other: "Activate",
	},
	{
		ID:    "admin",
		Other: "Admin",
	},
	{
		ID:    "forgot_password",
		Other: "Forgot Password",
	},
	{
		ID:    "forgot_password_success",
		Other: "An email with instructions describing how to reset your password has been sent.",
	},
	{
		ID:    "password_reset",
		Other: "Password Reset",
	},
	{
		ID:    "password_reset_email",
		Other: "Use the following link to reset your password. If this was not requested by you, please ignore this email.\n%s",
	},
	{
		ID:    "login",
		Other: "Login",
	},
	{
		ID:    "login_error",
		Other: "Could not login, please make sure that you have typed in the correct email and password. If you have forgotten your password, please click the forgot password link below.",
	},
	{
		ID:    "login_activated_error",
		Other: "Account is not activated yet.",
	},
	{
		ID:    "404_not_found",
		Other: "404 Not Found",
	},
	{
		ID:    "register",
		Other: "Register",
	},
	{
		ID:    "password_too_short_error",
		Other: "Your password must be 8 characters in length or longer",
	},
	{
		ID:    "register_error",
		Other: "Could not register, please make sure the details you have provided are correct and that you do not already have an existing account.",
	},
	{
		ID:    "register_success",
		Other: "Thank you for registering. An activation email has been sent with steps describing how to activate your account.",
	},
	{
		ID:    "user_activation",
		Other: "User Activation",
	},
	{
		ID:    "user_activation_email",
		Other: "Use the following link to activate your account. If this was not requested by you, please ignore this email.\n%s",
	},
	{
		ID:    "resend_activation_email_subject",
		Other: "Resend Activation Email",
	},
	{
		ID:    "resend_activation_email_success",
		Other: "A new activation email has been sent if the account exists and is not already activated. Please remember to check your spam inbox in case the email is not showing in your inbox.",
	},
	{
		ID:    "reset_password",
		Other: "Reset Password",
	},
	{
		ID:    "reset_password_error",
		Other: "Could not reset password, please try again",
	},
	{
		ID:    "password_reset_success",
		Other: "Your password has successfully been reset.",
	},
	{
		ID:    "search",
		Other: "Search",
	},
	{
		ID:    "search_results",
		Other: "Search Results",
	},
	{
		ID:    "no_results_found",
		Other: "No results found",
	},
	{
		ID:    "404_message_1",
		Other: "The page you're looking for could not be found.",
	},
	{
		ID:    "click_here",
		Other: "Click here",
	},
	{
		ID:    "404_message_2",
		Other: "to return to the main page.",
	},
	{
		ID:    "admin_dashboard",
		Other: "Admin Dashboard",
	},
	{
		ID:    "dashboard_message",
		Other: "You now have an authenticated session, feel free to log out using the link in the navbar above.",
	},
	{
		ID:    "footer_message_1",
		Other: "Fork this project on",
	},
	{
		ID:    "created_by",
		Other: "Created by",
	},
	{
		ID:    "forgot_password",
		Other: "Forgot password?",
	},
	{
		ID:    "forgot_password_message",
		Other: "Use the form below to reset your password. If we have an account with your email you will receive instructions on how to reset your password.",
	},
	{
		ID:    "email_address",
		Other: "Email address",
	},
	{
		ID:    "request_reset_email",
		Other: "Request reset email",
	},
	{
		ID:    "lang_key",
		Other: "en",
	},
	{
		ID:    "home",
		Other: "Home",
	},
	{
		ID:    "admin",
		Other: "Admin",
	},
	{
		ID:    "logout",
		Other: "Logout",
	},
	{
		ID:    "login",
		Other: "Login",
	},
	{
		ID:    "register",
		Other: "Register",
	},
	{
		ID:    "search",
		Other: "Search",
	},
	{
		ID:    "index_message_1",
		Other: "A simple website with user login and registration.",
	},
	{
		ID:    "index_message_2",
		Other: "The frontend uses",
	},
	{
		ID:    "index_message_3",
		Other: "and the backend is written in",
	},
	{
		ID:    "index_message_4",
		Other: "Read more about this project on",
	},
	{
		ID:    "password",
		Other: "Password",
	},
	{
		ID:    "login_terms",
		Other: "By pressing the button below to login you agree to the use of cookies on this website.",
	},
	{
		ID:    "request_new_activation_email",
		Other: "Request a new activation email",
	},
	{
		ID:    "resend_activation_email",
		Other: "Resend Activation Email",
	},
	{
		ID:    "resend_activation_email_message",
		Other: "If you have already registered but never activated your account you can use the form below to request a new activation email.",
	},
	{
		ID:    "request_activation_email",
		Other: "Request activation email",
	},
	{
		ID:    "reset_password_message",
		Other: "Please enter a new password.",
	},
	{
		ID:    "label_users",
		Other: "Users",
	},
	{
		ID:    "label_reports",
		Other: "Reports",
	},
	{
		ID:    "label_admin_panel",
		Other: "Admin Panel",
	},
	{
		ID:    "notifications_empty",
		Other: "Notifications empty",
	},
	{
		ID:    "messages_empty",
		Other: "Messages empty",
	},
	{
		ID:    "label_username",
		Other: "Username",
	},
	{
		ID:    "label_email",
		Other: "Email",
	},
	{
		ID:    "label_activated_at",
		Other: "Activated at",
	},
	{
		ID:    "label_member_since",
		Other: "Member since [days]",
	},
	{
		ID:    "label_role",
		Other: "Role",
	},
	{
		ID:    "label_status",
		Other: "Status",
	},
	{
		ID:    "label_no_reports",
		Other: "No. reports",
	},
	{
		ID:    "label_actions",
		Other: "Actions",
	},
	{
		ID:    "label_next",
		Other: "Next",
	},
	{
		ID:    "label_prev",
		Other: "Previous",
	},
	{
		ID:    "label_entries",
		Other: "entries",
	},
	{
		ID:    "label_user_reported",
		Other: "User reported",
	},
	{
		ID:    "label_user_reporting",
		Other: "User reporting",
	},
	{
		ID:    "label_accepted_by",
		Other: "Accepted by",
	},
	{
		ID:    "label_created_at",
		Other: "Created at",
	},
	{
		ID:    "label_starts_at",
		Other: "Starts at",
	},
	{
		ID:    "label_ends_at",
		Other: "Ends at",
	},
	{
		ID:    "label_days_to_unban",
		Other: "Days to unban",
	},
	{
		ID:    "label_showing",
		Other: "Showing",
	},
	{
		ID:    "label_no_records",
		Other: "No records to show",
	},
	{
		ID:    "label_to",
		Other: "to",
	},
	{
		ID:    "label_of",
		Other: "of",
	},
	{
		ID:    "status_active",
		Other: "Active",
	},
	{
		ID:    "status_pending",
		Other: "Pending",
	},
	{
		ID:    "status_banned",
		Other: "Banned",
	},
	{
		ID:    "profile",
		Other: "Profile",
	},
	{
		ID:    "settings",
		Other: "Settings",
	},
	{
		ID:    "sessions",
		Other: "Sessions",
	},
	{
		ID:    "report",
		Other: "Report",
	},
	{
		ID:    "accept",
		Other: "Accept",
	},
	{
		ID:    "welcome_index",
		Other: "Welcome to index page",
	},
	{
		ID:    "welcome_internal",
		Other: "Welcome to internal page, you are logged in as",
	},
	{
		ID:    "username_not_valid",
		Other: "Username not valid, you can use only letters, digits and underscore",
	},
	{
		ID:    "label_you_are_banned",
		Other: "You are banned, check out why ",
	},
	{
		ID:    "label_here",
		Other: "here",
	},

}
