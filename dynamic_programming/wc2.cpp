#include <iostream>
#include <string>
using namespace std;
class Solution {
public:
    bool isMatch(string input, string pattern) {
        // Get the lengths of the input and pattern strings
        int inputLen = input.length();
        int patternLen = pattern.length();

        // Initialize indices for the input and pattern strings
        int inputIndex = 0;
        int patternIndex = 0;

        // Initialize a flag to track if a match has been found
        bool matchFound = false;

        // Iterate through the input and pattern strings
        while (inputIndex < inputLen && patternIndex < patternLen) {
            // If the current characters match or pattern has '.', move both indices forward
            if (pattern[patternIndex] == '.' || pattern[patternIndex] == input[inputIndex]) {
                patternIndex++;
                inputIndex++;
            }
            // If the current character in the pattern is '*', try to match the remaining pattern with
            // the remaining input characters
            else if (pattern[patternIndex] == '*') {
                // Move the pattern index forward to skip the '*'
                patternIndex++;

                // If the '*' is the last character in the pattern, it matches the remaining input string
                if (patternIndex == patternLen) {
                    matchFound = true;
                    break;
                }

                // Recursively check if the remaining pattern matches the remaining input
                if (isMatch(input.substr(inputIndex), pattern.substr(patternIndex))) {
                    matchFound = true;
                    break;
                }
            }
            // If the characters do not match and the pattern does not have '*', return false
            else {
                matchFound = false;
                break;
            }
        }

        // If both indices have reached the end of their respective strings, the pattern is found
        if (inputIndex == inputLen && patternIndex == patternLen) {
            matchFound = true;
        }

        return matchFound;
    }
};


int main() {
    Solution s;
    cout << s.isMatch("aa", "a") << endl;
    cout << s.isMatch("aa", "*") << endl;
    cout << s.isMatch("cb", ".a") << endl;
    cout << s.isMatch("adceb", "*a*b") << endl;
    cout << s.isMatch("acdcb", "a*c.b") << endl;
    cout << s.isMatch("acdcb", "a*c.b*") << endl;
    cout << s.isMatch("acdcb", "a*c.b**") << endl;
    cout << s.isMatch("acdcb", "a*c.b***") << endl;
    cout << s.isMatch("acdcb", "a*c.b****") << endl;
    cout << s.isMatch("acdcb", "a*c.b*****") << endl;
    cout << s.isMatch("acdcb", "a*c.b******") << endl;
    cout << s.isMatch("acdcb", "a*c.b*******") << endl;
    cout << s.isMatch("acdcb", "a*c.b********") << endl;
    cout << s.isMatch("acdcb", "a*c.b*********") << endl;
    cout << s.isMatch("acdcb", "a*c.b**********") << endl;
    cout << s.isMatch("acdcb", "a*c.b***********") << endl;
    cout << s.isMatch("acdcb", "a*c.b************") << endl;
    cout << s.isMatch("acdcb", "a*c.b*************") << endl;
    cout << s.isMatch("acdcb", "a*c.b**************") << endl;
    cout << s.isMatch("acdcb", "a*c.b***************") << endl;
    cout << s.isMatch("acdcb", "a*c.b****************") << endl;
    cout << s.isMatch("acdcb", "a*c.b*****************") << endl;
    cout << s.isMatch("acdcb", "a*c.b******************") << endl;
    cout << s.isMatch("acdcb", "a*c.b*******************") << endl;
    cout << s.isMatch("acdcb", ".*") << endl;

    return 0;
}
