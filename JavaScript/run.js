
const prompt = require('prompt-sync')();
var game = require('./asynchro.js');
var turn = 1;
var end = false;

const rl = game.rl;

rl.on('close', () => {
    console.log('Exiting...');
    process.exit(0);
});

console.log('Welcome to Jarnac');
let player1 = {
    name: prompt('Enregistring your name as player 1 : '),
    plate: [[[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []]],
    hand: [],
    words_played: []
};
let player2 = {
    name: prompt('Enregistring your name as player 2 : '),
    plate: [[[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []]],
    hand: [],
    words_played: []
}
let players = [player1, player2];

// 异步获取初始手牌
const initHandsPromises = [
    game.get_random_letters(true),
    game.get_random_letters(true)
];

Promise.all(initHandsPromises)
    .then(async initHands => {
        for (let letter of initHands[0]) {
            player1.hand.push(letter);
        }
        for (let letter of initHands[1]) {
            player2.hand.push(letter);
        }
        console.log('Game start');
        await first_turn();
        await play_turns();
    });




    
async function first_turn() {
    console.log(`************* Turn of player ${turn} *************`);
    let player = players[turn - 1];
    game.display_plate(player);
    game.display_hand(player);
    console.log('_____________________________________________________');

    let play_type = 2;
    let row = 0;
    let letters = await game.get_aimed_letter(player, play_type);
    console.log('you chose:', letters);
    let word;
    if (play_type === 1) {
        word = await game.rearrange_letters(player.words_played[row].concat(letters));
        while (!game.verify_word(word)) {
            console.log('Invalid word. Please try again.');
            word = await game.rearrange_letters(letters);
        }
    } else {
        word = letters;
        while (!game.verify_word(word)) {
            console.log('Invalid word. Please try again.');
            for (let letter of word) {
                player.hand.push(letter);
            }
            console.log(player.hand);
            word = await game.get_aimed_letter(player, 2);
        }
    }
    game.put_word(player, word, row);
    game.display_plate(player);
    game.display_hand(player);
    turn = 2;
}



async function play_turns() {
    console.log(`************* Turn of player ${turn} *************`);
    let player = players[turn - 1];
    game.display_plate(player);
    game.display_hand(player);
    console.log('_________________________________________________________');
    let action = await game.ask('choose an action: 1- exchange letters, 2- put a word: ');
    while (action !== '1' && action !== '2') {
        action = await game.ask('please enter a correct number 1/2: ');
    }
    if (action === '1') {
        let to_exchange = await game.get_aimed_letter(player, 'exchange');
        console.log(to_exchange);
        game.exchange_with_bag(player, to_exchange);
    }
    game.display_plate(player);
    game.display_hand(player);
    let res = await game.get_aimed_row(player);
    let play_type = res[0];
    let row = res[1];
    let letters = await game.get_aimed_letter(player, play_type);
    console.log('you chose:', letters);
    let word;
    if (play_type === 1) {
        word = await game.rearrange_letters(player.words_played[row].concat(letters));
        while (!game.verify_word(word)) {
            console.log('Invalid word. Please try again.');
            word = await game.rearrange_letters(letters);
        }
    } else {
        word = letters;
        while (!game.verify_word(word)) {
            console.log('Invalid word. Please try again.');
            for (let letter of word) {
                player.hand.push(letter);
            }
            console.log(player.hand);
            word = await game.get_aimed_letter(player, 2);
        }
    }
    game.put_word(player, word, row);
    game.display_plate(player);
    game.display_hand(player);
    turn = (turn === 1) ? 2 : 1;
    if (!end) {
        await play_turns();
    }
}