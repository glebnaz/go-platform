package engine

import log "github.com/sirupsen/logrus"

const (
	banner = `
===================================================================⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀Witcher⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
===================================================================⠀
⠀⠀⠀⠀⠀⠀⢠⣾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣳⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣿⣿⣶⣄⣀⠀⠀⠀⠀⠀⠀⠈⢳⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⡞⠁⠀⠀⠀⠀⠀⠀⣀⣠⣶⣿⣿⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠘⠷⣽⣻⢿⣿⣶⣤⣄⡀⠀⠀⠀⠹⣿⣷⣄⢶⣄⠹⣿⣿⣿⣿⣿⣿⣿⣿⠏⣠⡶⣠⣾⣿⠏⠀⠀⠀⢀⣠⣤⣶⣿⡿⣟⣯⡾⠃⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢱⠈⠛⢷⣮⣟⠿⣿⣿⣿⣶⣦⣄⡘⢿⣿⡎⣿⣷⡹⣿⣿⣿⣿⣿⣿⢏⣾⡿⢱⣿⡿⢃⣠⣴⣶⣿⣿⣿⠿⣻⣽⡾⠛⠁⡎⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢣⢳⣄⡉⠻⢿⣶⣝⡻⢿⣿⢟⣼⣷⣿⣿⣿⣿⣷⣽⣿⡿⢿⣿⣯⣾⣿⣿⣿⣿⣿⣧⡻⣿⡿⢟⣫⣶⡿⠟⢉⣠⡞⡼⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⢧⡹⣿⣦⣀⠈⠛⢿⡷⣠⣛⠿⣿⣿⣏⣍⢿⣿⣿⡇⣿⣿⢸⣿⣾⡿⣩⣹⣿⣿⠿⣛⣄⢾⡿⠛⠁⣀⣴⣿⢏⡼⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠳⡜⢿⣿⣷⡄⠀⢘⠻⢿⣿⣮⣝⡺⢿⣷⡽⣿⣿⢻⡟⣿⣿⢯⣾⡿⢗⣫⣵⣿⡿⠟⡃⠀⢠⣾⣿⡿⢣⠞⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠈⠉⠛⠻⠿⢿⣆⠙⢦⣙⠿⢀⣰⣿⣷⡀⠈⣝⡻⢿⣷⣯⣓⢹⣿⢸⡇⣿⡏⣚⣽⣾⡿⢟⣫⠁⢀⣾⣿⣆⡀⠿⣋⡴⠋⣰⡿⠿⠟⠛⠉⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣄⡺⢿⣿⣿⣟⣷⣄⠘⠛⠛⠈⠛⠻⢼⣿⡟⢳⣿⡧⠿⠛⠁⠛⠛⠃⣀⣾⣹⣿⣿⡿⢗⣠⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣾⣿⣿⣿⣿⣷⣬⣛⠿⣮⣝⣻⣿⣓⡍⣷⢸⣿⣿⡇⢸⣿⣿⡇⣶⢩⣚⣿⣟⣫⣵⠿⣛⣥⣶⣿⣿⣿⣿⣷⣶⣤⣀⡀⠀⠀⠀⠀⠀⠀
⠀⢀⣀⣤⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣝⠻⢿⣿⣿⡿⣼⣿⣿⣿⣿⣿⣿⣧⢿⣿⣿⡿⠟⣫⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣤⣀⡀⠀
⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠉⢻⣝⡻⢿⣿⣿⣿⣿⡿⢟⣫⡟⠉⠀⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣴⣾⣿⣿⣿⣿⡶⠄⠸⣿⣿⣄⡀⠀⠀⢀⣠⣿⣿⡇⠠⢶⣿⣿⣿⣿⣷⣦⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⣿⣿⣿⣿⡿⠟⠉⠀⠀⠀⠛⠛⠛⠛⢛⡛⠛⠛⠛⠛⠀⠀⠀⠉⠻⢿⣿⣿⣿⣿⣿⣿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⢸⣧⢠⡀⠀⠀⢀⡄⣼⡇⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣀⡀⠀⢸⠃⠸⠁⠏⠹⠈⠇⠘⡇⠀⢀⣀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣀⣴⣿⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣦⣀⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣾⣿⣿⠟⠋⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⠏⠀⠀⠀⣿⡄⢀⠀⠀⠀⠀⡀⢠⣿⠀⠀⠀⠹⣿⣿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠙⠻⣿⣿⣷⡄⠀⠀⠀⠀
⠀⠀⠀⣰⣿⠟⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⠃⠀⠀⠀⠀⢹⠇⢸⣷⣄⣠⣾⡇⠸⡟⠀⠀⠀⠀⠘⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠻⣿⣆⠀⠀⠀
⠀⠀⠞⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⡿⠁⠀⠀⠀⠀⠀⢸⢸⢸⣿⣿⣿⣿⡇⡇⡇⠀⠀⠀⠀⠀⠈⢿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠳⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⡟⠀⠀⠀⠀⠀⠀⠀⢸⡜⠰⣿⣿⣿⣿⠆⢣⡇⠀⠀⠀⠀⠀⠀⠀⢻⣿⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⠈⠁⠈⠁⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⡄⠀⠀⠀⠀⢠⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢃⡷⢠⣤⣤⡄⢾⡸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢨⣥⣬⣉⣉⣥⣬⡅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠿⢿⣿⣿⡿⠿⠃
===================================================================⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀High performance, app-engine!⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
===================================================================⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`

	introduce = `Server is running! Your Debug Port is: %s`
)

func (s *Server) introduce() {
	log.Infof("Server Is Ready")
	log.Infof(introduce, s.PORT)
	if !s.disableBanner {
		log.Infof(banner)
	}
}
