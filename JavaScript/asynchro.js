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
    let letters = [];
    let answer = await ask(`Which letters do you want to ${type} ? `);
    for (let letter of answer) {
        letters.push(letter.toUpperCase());
    }
    if (['P','A','S','S'].every(letter => letters.includes(letter))) {
        return 'PASS';
    }
    if (type ==='exchange') {
        if (letters.length !== 3) {
            console.log('You must exchange 3 letters');
            return get_aimed_letter(player, 'exchange');
        }
    }
    if (!Object.entries(countLetters(letters)).every(([letter, count]) => {
        return countLetters(player.hand)[letter] >= count;
    })) {
        console.log(`You do not have the letter in your hand`);
        return get_aimed_letter(player, 'exchange');
    }
    for (let letter of letters) {    
        player.hand.splice(player.hand.indexOf(letter), 1);
    }    
    return letters;
}
function countLetters(word) {
    const count = {};
    for (let i = 0; i < word.length; i++) {
        const letter = word[i];
        if (!count[letter]) {
            count[letter] = 0;
        }
        count[letter]++;
    }
    return count;
}

async function rearrange_letters(letters) {
    let arranged_letters = [];
    console.log('Enter all the letters in the row in your desired order: ', letters);
    console.log('Enter nothing to go back to the previous step.');
    const input = await ask('Letters (separated by spaces): ');
    if (input === '') {
        return [''];
    }
    const inputLetters = input.split(' ');

    for (const letter of inputLetters) {
        const index = letters.indexOf(letter);

        if (index !== -1) {
            letters.splice(index, 1);
            arranged_letters.push(letter);
        } else {
            console.log('Invalid letter. Please try again.');
        }
    }

    return arranged_letters;

}

/*
 * check if the word is valid, return true or false
 * ! problem with the accent
 */
function verify_word(letters) {
    let word = letters.join('');
    if (words.includes(word)) {
        return true;
    } else {
        return false;
    }
}


/**
 * put the verified word in the plate so there is no verification here
 * @param {*} player 
 * @param {*} word verified word
 * @param {*} x row number
 */
function put_word(player, word, x) {
    let row = player.plate[x];
    let i = 0;
    let played = [];
    for (let letter of word) {
        row[i].push(letter);
        played.push(letter);
        i++;
    }
    player.words_played.push(played);

}


/*
 * display the plate of the player
 * return: nothing
 */
function display_plate(player) {
    for (let row of player.plate) {
      let rowStr = '';
      for (let cell of row) {
        rowStr += (Array.isArray(cell) && cell.length === 0) ? ' [ ] ' : ' [' + cell + '] ';
      }
      console.log(rowStr);
    }
  }
  
  /*
   * display the hand of the player
   * return: nothing
   */
  function display_hand(player) {
    console.log('hand of player ', player.name, ': ', player.hand);
  }
  
  module.exports = {
    rl,
    ask,
    get_random_letters,
    exchange_with_bag,
    get_aimed_row,
    get_aimed_letter,
    rearrange_letters,
    verify_word,
    put_word,
    display_plate,
    display_hand
  };