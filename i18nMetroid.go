package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var i18nbundle *i18n.Bundle

func init() {
	i18nbundle = i18n.NewBundle(language.English)
	i18nbundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	i18nbundle.MustParseMessageFileBytes([]byte(`
HelloWorld = "Bonjour !"
CouldNotReadRequest = "Nous ne pouvons pas traiter ta demande."
InvalidTokenPassword = "Ce lien est expiré. As-tu reçu un lien plus récent ?"
EmailTokenMismatch = "Email d'authentification expiré. Réessayer de vous inscrire depuis le début."
InternalError = "Erreur interne."
EmailNicknameValidation = "L'email et le pseudo doivent contenir entre 3 et 12 caractères."
EmailAlreadyTaken = "Ton enregistrement a déjà été pris en compte "
NicknameAlreadyTaken = "Ce pseudo est déjà utilisé 😥"
EmailDoesNotExist = "Cet email n'existe pas 😥"
InvalidEmail = "Email non valide 😥"
InvalidNickname = "Pseudo non valide 😥"
NickNameValidation = "Ton pseudo ne peut pas commencer avec un chiffre ou contenir des espaces ou des caractères spéciaux."
ThankYouWelcome = "Merci de t'être enregistré sur notre liste de diffusion !"
notWhitelisted = "Tu n'as pas encore reçu d'invitation. Contacte-nous à contact@piepacker.com pour obtenir l'accès."
passwordReset = "Email de réinitialisation de mot de passe envoyé pour "
ksBackerSignup = "Merci de t'être enregistré"
ksBackerSecondPart = "Nous t'avons envoyé un email pour te demander de confirmer ton inscription"
`), "fr.toml")
	i18nbundle.MustParseMessageFileBytes([]byte(`
HelloWorld = "Hello World!"
CouldNotReadRequest = "We cannot process your request."
InvalidTokenPassword = "This password link is expired. Have you received a more recent link?"
EmailTokenMismatch = "Expired email link. Please try again from to signup from the beginning."
InternalError = "Internal error."
EmailNicknameValidation = "Email and Nickname must contain between 3 and 12 characters."
EmailAlreadyTaken = "You are already registered"
NicknameAlreadyTaken = "This nickname is already taken 😢"
EmailDoesNotExist = "Email does not exist 😢"
InvalidEmail = "Invalid email 😢"
InvalidNickname = "Invalid nickname 😢"
notWhitelisted = "You have not yet been invited. Contact us at contact@piepacker.com to get access."
NickNameValidation = "Nickname cannot start with a number or contain spaces or symbols."
ThankYouWelcome = "Thank you for joining our mailing list!"
passwordReset = "Sent password reset email for "
ksBackerSignup = "Thank you for subscribing"
ksBackerSecondPart = "We have sent you an email to ask you confirm your subscription"
`), "en.toml")
	i18nbundle.MustParseMessageFileBytes([]byte(`
HelloWorld = "¡Hola mundo!"
CouldNotReadRequest = "No podemos procesar su solicitud"
InvalidTokenPassword = "Este link de contraseña ha expirado. ¿Tienes un link más reciente?"
EmailTokenMismatch = "Enlace de correo electrónico caducado. Vuelve a intentarlo desde para registrarte desde el principio."
InternalError = "Error interno"
EmailNicknameValidation = "El email y nombre de usuario debe contener entre 3 y 12 caracteres"
EmailAlreadyTaken = "Ya estás registrado"
NicknameAlreadyTaken = "Este nombre de usuario ya está en uso 😢"
EmailDoesNotExist = "El email no existe"
InvalidEmail = "email inválido"
InvalidNickname = "Nombre de usuario inválido"
notWhitelisted = "Todavía no has sido invitado. Contáctanos a través de contact@piepacker.com para conseguir acceso."
NickNameValidation = "El nombre de usuario no puede empezar con un número, ni contener espacios o símbolos."
ThankYouWelcome = "¡Gracias por unirte a nuestra lista de emails!"
passwordReset = "Enviado un correo electrónico de restablecimiento de contraseña para "
ksBackerSignup = "Gracias por suscribirse"
ksBackerSecondPart = "Le hemos enviado un correo electrónico para pedirle que confirme su suscripción"
`), "es.toml")
}

func getText(localizer *i18n.Localizer, msg string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: msg})
}

func translateInternalError(localizer *i18n.Localizer, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %s", getText(localizer, "InternalError"), err.Error())
}
