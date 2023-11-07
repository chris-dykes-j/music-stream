function selectTrack(track) {
    const source = document.getElementById('player-track')
    source.setAttribute("src", `http://localhost:8080/music/${track}`) 
    const player = document.getElementById('player')
    player.load()
    player.play()
}
