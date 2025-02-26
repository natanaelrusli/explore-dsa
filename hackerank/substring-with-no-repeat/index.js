/**
 * 
 * @param {*} s 
 * @returns length of the longest substring without repeating characters
 * 
 * Time Complexity: O(n)
 * Uses the slidng window algorithm
 */

function longestSubstringWithNoRepeat(s) {
  let start = 0; // Start of the current window
  let maxLength = 0; // Maximum length of substring without repeats
  let charIndexMap = {}; // Map to store the latest index of each character

  for (let end = 0; end < s.length; end++) {
    const char = s[end];

    // If character is already in the window, update the start position
    if (char in charIndexMap && charIndexMap[char] >= start) {
      start = charIndexMap[char] + 1;
    }

    // Update the character's latest index
    charIndexMap[char] = end;
    // Calculate the maximum length
    maxLength = Math.max(maxLength, end - start + 1);
  }

  return maxLength;
}

console.log(longestSubstringWithNoRepeat("abcabcbb")); // Output: 3 (substring "abc")
