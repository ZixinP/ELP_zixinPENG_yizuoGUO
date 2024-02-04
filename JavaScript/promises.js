const { get } = require('http');
const readline = require('readline');


const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});


function ask(question) {
    return new Promise(resolve => {
        rl.question(question, answer => {
            resolve(answer);
        });
    });
}

async function get_aimed_letters(type){
    let letters = [];
    let answer = await ask(`Which letters do you want to ${type} ? `);
    for (let letter of answer) {
        letters.push(letter.toUpperCase());
    }
    return letters;
}
get_aimed_letters('exchange').then(values => {
    console.log(values);
    rl.close();
});


// Promise.then(get_aimed_letters('exchange')).then(values => {
//     console.log(values);
//     if (values === 'PASS') {
//         console.log('you passed your turn');
//         // turn = (turn === 1) ? 2 : 1;
//         // play_turns();
//     } else {
//         // for (let letter of values) {
//         //     player.hand.push(get_random_letters(false)[0]);
//         //     if (letters_left[letter] === undefined) {
//         //         letters_left[letter] = 0;
//         //     }
//         //     letters_left[letter]++;
//         // }
//         console.log('you chose:', values);s
//     }
// });