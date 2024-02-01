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
    hand : [],
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
    hand : [],
    words_played : []
    }
let players = [player1, player2];
let init_hands = [game.get_random_letters(true),game.get_random_letters(true)];
for (let letter of init_hands[0]) {
    player1.hand.push(letter);
}
for (let letter of init_hands[1]) {
    player2.hand.push(letter);
}
console.log('Game start');
while (!end) {
    console.log('Turn of player' + turn);
    let player = players[turn - 1];
    game.display_plate(player);
    game.display_hand(player);
    
    console.log('_________________________________________________________');
    
    let action = prompt('choose an action : 1- exchange letters, 2- put a word: ');
    while (action !== '1' && action !== '2'){
        action = prompt('please enter a correct number 1/2 : ');
    }
    if (action === '1') {
        let to_exchange = game.get_aimed_letter(player,'exchange')
        console.log(to_exchange);
        game.exchange_with_bag(player,to_exchange);
    } 
    game.display_plate(player);
    game.display_hand(player);
    let res = game.get_aimed_row(player);
    let play_type = res[0];
    let row = res[1];
    let letters = game.get_aimed_letter(player,play_type);
    console.log('you chose:', letters)
    let word = game.rearrange_letters(letters);
    while (!game.verify_word(word)) {
        console.log('Invalid word. Please try again.');
        word = game.rearrange_letters(letters);
    }
    game.put_word(player,word,row);
    game.display_plate(player);
    game.display_hand(player);
    turn = (turn === 1) ? 2 : 1;
} 