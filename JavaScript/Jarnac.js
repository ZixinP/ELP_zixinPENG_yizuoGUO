//适用于node.js的jarnac游戏
const prompt = require('prompt-sync')();

const letter_pool = {"A":14,"B":4,"C":7,"D":5,"E":19,"F":2,"G":4,"H":2,"I":11,"J":1,"K":1,"L":6,"M":5,"N":9,"O":8,"P":4,"Q":1,"R":10,"S":7,"T":9,"U":8,"V":2,"W":1,"X":1,"Y":1,"Z":2};
let letter_left = letter_pool;
// var init = false;

function get_random_letters(init) {
  const letters = Object.keys(letter_left);
  const random_letters = [];
  let nb_letters = 1;
  if (init) { nb_letters = 6; } 

  for (let i = 0; i < nb_letters ; i++) {
    let letter = '';
    
    while (letter === '' || letter_left[letter] === 0) {
      const random_index = Math.floor(Math.random() * letters.length);
      letter = letters[random_index];
    }
    
    random_letters.push(letter);
    letter_left[letter]--;
  }
  
  return random_letters;
}


// function for test: count the number of letters left
function count_letters_left() {
  let count = 0;
  for (let letter in letter_left) {
    count += letter_left[letter];
  }
  return count;
}

// for (let i = 0; i < 5; i++) {
//   console.log(count_letters_left());
//   console.log(get_random_letters());
// }



let player1 = {
  plate : [[],[],[],[],[],[],[],[]],
  hand : [],
  words_played : []
}
let player2 = {
  plate : [[[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]],
           [[],[],[],[],[],[],[],[],[]]
          ],
  hand : [],
  words_played : []
}

function letter_aimed(type){
  if (type === 'put') {
    let letter = prompt('choose a letter to put:');
    if (player1.hand.includes(letter)) {
      return letter;
    } else {
      console.log('you do not have this letter in your hand, try again');
      return letter_aimed('put');
    }
  }
  else if (type === 'exchange') {
    return [letter_aimed('put'),letter_aimed('put'),letter_aimed('put')];
  }
}

function put_letter(letter, player, x, y) {
  x = parseInt(x);
  y = parseInt(y);
  if (Array.isArray(player.plate[x][y]) && player.plate[x][y].length > 0) {
    console.log('this cell is already occupied,tried another one');
    let newLetter = prompt('letter:');
    let newX = parseInt(prompt('x:'));
    let newY = parseInt(prompt('y:'));
    put_letter(newLetter, player, newX, newY);
  } else {
    player.plate[x][y] = [letter];
  }
}

function display_plate(player) {
  for (let row of player.plate) {
    let rowStr = '';
    for (let cell of row) {
      rowStr += (Array.isArray(cell) && cell.length === 0) ? ' [ ] ' :  ' ['+cell+'] ';
    }
    console.log(rowStr);
  }
}
put_letter('A', player2, 0, 0);
display_plate(player2);
put_letter('B', player2, 0, 0);
display_plate(player2);
