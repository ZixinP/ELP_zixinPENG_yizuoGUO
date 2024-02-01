const prompt = require('prompt-sync')();
var game = require('./Jarnac.js');
var turn = 1;
var end = false;

console.log('Welcome to Jarnac');
let player1 = {
    name : prompt('Enregistring your name as player 1 : '),
    plate : [[[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]]],
    hand : [game.get_random_letters(true)],
    words_played : []
    };
let player2 = {
    name : prompt('Enregistring your name as player 2 : '),
    plate : [[[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]],
             [[],[],[],[],[],[],[],[],[]]],
    hand : [game.get_random_letters(true)],
    words_played : []
    }
let players = [player1, player2];
console.log('Game start');
while (!end) {
    console.log('Turn of player' + turn);
    let player = players[turn - 1];
    game.display_plate(player);
    game.display_hand(player);
}