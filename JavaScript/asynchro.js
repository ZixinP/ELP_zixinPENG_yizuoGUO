const readline = require('readline');
var words = require('an-array-of-french-words').map(word => word.toUpperCase());
/*https://github.com/words/an-array-of-english-words*/

const letter_pool = { "A": 14, "B": 4, "C": 7, "D": 5, "E": 19, "F": 2, "G": 4, "H": 2, "I": 11, "J": 1, "K": 1, "L": 6, "M": 5, "N": 9, "O": 8, "P": 4, "Q": 1, "R": 10, "S": 7, "T": 9, "U": 8, "V": 2, "W": 1, "X": 1, "Y": 1, "Z": 2 };
let letters_left = letter_pool;


const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

/*
 * get 6 random letters from the bag if it concerns the first hand of the game or 1 letter if not
 * return: an array of 6 letters or a letter
 */
function get_random_letters(init) {
    const letters = Object.keys(letters_left);
    const random_letters = [];
    let nb_letters = 1;
    if (init) { nb_letters = 6; }

    for (let i = 0; i < nb_letters; i++) {
        let letter = '';

        while (letter === '' || letters_left[letter] === 0) {
            const random_index = Math.floor(Math.random() * letters.length);
            letter = letters[random_index];
        }

        random_letters.push(letter);
        letters_left[letter]--;
    }

    return random_letters;
}

/*
 * change the letters chosen with the bag, no return
 */
function exchange_with_bag(player, letters) {
    for (let letter of letters) {
        player.hand.push(get_random_letters(false)[0]);
        if (letters_left[letter] === undefined) {
            letters_left[letter] = 0;
        }
        letters_left[letter]++;
    }
}

function ask(question) {
    return new Promise(resolve => {
        rl.question(question, answer => {
            resolve(answer);
        });
    });
}

rl.on('close', () => {
    console.log('Exiting...');
    process.exit(0);
});

async function get_aimed_row(player) {
    console.log('choose a row (0-7) to put the letter');

    const x = await ask('x: ');
    let parsed_x = parseInt(x);

    while (isNaN(parsed_x)) {
        console.log('Invalid input. Please try again.');
        const newX = await ask('x: ');
        parsed_x = parseInt(newX);
    }

    if (parsed_x < 0 || parsed_x > 7) {
        console.log('this row is outside the plate, choose again');
        return get_aimed_row(player);
    } else {
        if (player.plate[parsed_x][0].length === 0) {
            let aime = parsed_x;
            for (let i = parsed_x; i >= 0; i--) {
                if (player.plate[i][0].length === 0) {
                    aime = i;
                }
            }
            console.log('you try to put a letter in an empty row, you have to fill the row from the top');
            console.log('please choose at least 3 letters in your hand to put in the row ', aime);
            return [2, aime];
        } else {
            console.log('you try to put a letter in a row with letters, please choose a letter in your hand to put in this row');
            return [1, parsed_x];
        }
    }
}

async function get_aimed_letter(player, type) {
    if (type === 1) {
        const letter = await ask('choose a letter: ');

        if (player.hand.includes(letter)) {
            const index = player.hand.indexOf(letter);
            player.hand.splice(index, 1);
            return letter;
        } else {
            console.log('you do not have this letter in your hand, try again');
            return get_aimed_letter(player, 1);
        }
    } else if (type === 2 || type === 'exchange') {
        console.log(`choose 3 letters to ${type === 2 ? 'put' : 'exchange'}`);
        const letter1 = await get_aimed_letter(player, 1);
        const letter2 = await get_aimed_letter(player, 1);
        const letter3 = await get_aimed_letter(player, 1);

        return [letter1, letter2, letter3];
    }
}
let player1 = {
    name: 'Player 1',
    plate: [[['D'], ['O'], ['G'], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []],
    [[], [], [], [], [], [], [], [], []]],
    hand: ['E', 'R', 'T', 'S', 'A', 'I'],
    words_played: []
};

get_aimed_letter(player1, 2);