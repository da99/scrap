
export function constant_time_compare(a : string, b : string) {
  let result = 0;
  let a_length = a.length;
  for(let i = 0; i < a_length; i++) {
    result |= a.charCodeAt(i) ^ b.charCodeAt(i);
  }
  // Leave length comparison for last
  // or else attacker will know the length of target.
  result |= a.length ^ b.length;
  return result === 0;
}; // function

