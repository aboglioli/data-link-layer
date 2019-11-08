package physical

// Interface de comunicación.
// Servicios prestados por la capa física.
// Se envían bytes desde las capas físicas de emisor y receptor.
// Se han implementado tanto con puerto serial como a través de sockets TCPs. Es
// indistinto cuál se utilice ya que para este ejemplo son simples "tuberías"
// para comunicar punto a punto. En este ejemplo se envián y reciben bytes y es
// todo lo que la capa física ve.
type Interface interface {
	Send([]byte) error
	Recv() ([]byte, error)
}
