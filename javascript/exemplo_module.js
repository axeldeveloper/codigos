
/***   
 * user.js 
 * 
 */


const getName = () => {
    return 'Jim';
};
  
const getLocation = () => {
    return 'Munich';
};

  
const dateOfBirth = '12.01.1982';

exports.getName = getName;
exports.getLocation = getLocation;
exports.dob = dateOfBirth;


// demostrate
const user = require('./user');

console.log(
  `${user.getName()} lives in ${user.getLocation()} and was born on ${user.dob}.`
);





//or 

exports.getName = () => {
    return 'Jim';
};
  
exports.getLocation = () => {
    return 'Munich';
};
  
exports.dob = '12.01.1982';


//usage

const { getName, dob } = require('./user');
console.log(
  `${getName()} was born on ${dob}.`
);




