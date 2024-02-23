func (g *Game) RouteMessage() {
	log.Println("Starting RouteMessage ...")
BREAK:
	for {
		select {
		case msgBytes := <-g.Player1.MsgFromClient:
			msg := models.ParseSocketMessage(msgBytes)
			g.handleGameMove(msg, g.Player1, g.Player2)

		case msgBytes := <-g.Player2.MsgFromClient:
			msg := models.ParseSocketMessage(msgBytes)
			g.handleGameMove(msg, g.Player2, g.Player1)

		// Handle exit messages
		case <-g.Player1.ExitChan:
			log.Println(g.Player1.name, " has exited the game")
			g.Player2.SendMessage(getGameEndMessage(g.Player2.name))
			go g.Player2.Close()
			break BREAK

		case <-g.Player2.ExitChan:
			log.Println(g.Player2.name, " has exited the game")
			g.Player1.SendMessage(getGameEndMessage(g.Player1.name))
			go g.Player1.Close()
			break BREAK
		}
	}
	log.Println("Closing RouteMessage game for", g.Player1.name, " and ", g.Player2.name)
}