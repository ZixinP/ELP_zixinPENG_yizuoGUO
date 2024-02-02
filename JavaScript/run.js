const readline = require('readline');
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

        // 异步执行游戏循环
        async function playGame() {
            let turn = 1;
            let end = false;

            while (!end) {
                console.log(`************* Turn of player ${turn} *************`);

                let player = players[turn - 1];

                // 异步显示板和手牌
                await game.display_plate(player);
                await game.display_hand(player);

                console.log('_________________________________________________________');

                let action = await game.ask('choose an action: 1- exchange letters, 2- put a word: ');

                while (action !== '1' && action !== '2') {
                    action = await game.ask('please enter a correct number 1/2: ');
                }

                if (action === '1') {
                    let to_exchange = await game.get_aimed_letter(player, 'exchange');
                    console.log(to_exchange);
                    await game.exchange_with_bag(player, to_exchange);
                }

                // 异步显示板和手牌
                await game.display_plate(player);
                await game.display_hand(player);

                let res = await game.get_aimed_row(player);
                let play_type = res[0];
                let row = res[1];
                let letters = await game.get_aimed_letter(player, play_type);
                console.log('you chose:', letters);

                let word;

                if (play_type === 1) {
                    // 异步重新排列字母
                    word = await game.rearrange_letters(letters);

                    while (!await game.verify_word(word)) {
                        console.log('Invalid word. Please try again.');
                        word = await game.rearrange_letters(letters);
                    }
                } else {
                    word = letters;

                    while (!await game.verify_word(word)) {
                        console.log('Invalid word. Please try again.');

                        for (let letter of letters) {
                            player.hand.push(letter);
                        }

                        console.log(player.hand);
                        word = await game.get_aimed_letter(player, 2);
                    }
                }

                await game.put_word(player, word, row);

                // 异步显示板和手牌
                await game.display_plate(player);
                await game.display_hand(player);

                turn = (turn === 1) ? 2 : 1;
            }
        }

        // 调用异步游戏循环
        await playGame();
    });



// 在程序结束时关闭 readline


