package asciiart

func Split(str string) []string {
	slice := []string{}
	newStr := ""

	for i := 0; i < len(str); i++ {
		// verifie que \ suit d'un n.
		if i < len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			// si oui on verifie si il y'a qlq chose avant \n
			if newStr != "" {
				slice = append(slice, newStr)
				newStr = ""
			}
			// à la place de /n on ajoute un "" pour eviter le probleme de la duplication d \n
			slice = append(slice, "")
			// ici on saute l'element suivant qui est le n
			i += 1

		} else { // sinon on ajoute directement au string
			newStr += string(str[i])
		}
	}

	if newStr != "" {
		slice = append(slice, newStr)
	}
	return slice
}
