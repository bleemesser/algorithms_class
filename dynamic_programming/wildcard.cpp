// I've relearned c++ so I'll be using it for this project
#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <chrono>
#include <stdlib.h>

// Rather than the exact problem found here https://www.techiedelight.com/wildcard-pattern-matching/ 
// I'm choosing to do a different problem that checks whether the pattern is present somewhere within the string not 
// an exact match to the complete string
// this is more useful as a find function for a text editor or something, and I'll return the start and end indices of the pattern if it is found

using namespace std;

// naive solution 
bool match_wildcard_naive(string input, string pattern, int& found_pattern_start, int& found_pattern_end) {
    // Get the lengths of the input and pattern strings
    int inputLen = input.length();
    int patternLen = pattern.length();
    
    // Set the found pattern start and end to -1
    found_pattern_start = -1;
    found_pattern_end = -1;

    // Initialize indices for the input and pattern strings
    int inputIndex = 0;
    int patternIndex = 0;
    int junk = 0;

    // Iterate through the input and pattern strings
    while (inputIndex < inputLen && patternIndex < patternLen) {
        // If the current characters match or pattern has '?', move both indices forward
        if (pattern[patternIndex] == '?' || pattern[patternIndex] == input[inputIndex]) {
            // cout << "char match at " << inputIndex << endl;
            patternIndex++;
            inputIndex++;
        }
        // If the current character in the pattern is '*', try to match the remaining pattern with
        // all possible substrings of the input string
        else if (pattern[patternIndex] == '*') {
            // cout << "wildcard found at " << inputIndex << endl;
            // Save the current positions of the input and pattern indices
            int inputBackup = inputIndex;
            // Move the pattern index forward to skip the '*'
            patternIndex++;
            int patternBackup = patternIndex;

            // If the '*' is the last character in the pattern, it matches the remaining input string
            if (patternIndex == patternLen) {
                found_pattern_start = junk;
                found_pattern_end = inputLen - 1;
                return true;
            }
            // Loop through the input string from the current position
            while (inputIndex < inputLen) {
                // Recursively check if the remaining pattern matches the remaining input
                if (match_wildcard_naive(input.substr(inputIndex), pattern.substr(patternIndex), found_pattern_start, found_pattern_end)) {
                    found_pattern_start = junk;
                    found_pattern_end += inputIndex; // the recursive call returns the end index relative to the start of
                    // the recursive call, so we need to add the current index to get the end index relative to the start of the original call
                    // cout << "pattern (wc) found at " << inputIndex << endl;
                    return true;
                }
                // Move the input index forward to consider the next character
                inputIndex++;
            }

            // If no match is found, restore the input and pattern indices to the backup positions
            inputIndex = inputBackup;
            patternIndex = patternBackup;
        }
        // If the characters do not match and the pattern does not have '*', return false
        else {
            // cout << "char mismatch at " << inputIndex << endl;
            inputIndex++;
            junk++; // junk is used to keep track of the number of leading characters that don't match
        }
    }

    // If both indices have reached the end of their respective strings, the pattern is found
    if (inputIndex == inputLen && patternIndex == patternLen) {
        found_pattern_start = junk;
        found_pattern_end = inputLen - 1;
        // cout << "full pattern match" << endl;
        return true;
    }
    // If only the input string has reached the end, but the pattern has not, return false
    if (inputIndex == inputLen && pattern[patternIndex] != '*')
        return false;

    found_pattern_start = junk;
    found_pattern_end = inputIndex - 1;
    // cout << "pattern found at " << inputIndex << endl;
    return true;
}

// dynamic programming solution
bool match_wildcard_dp(string input, string pattern) {
    // get length of input and pattern strings
    int inputLen = input.length();
    int patternLen = pattern.length();

    // initialize the dp table
    vector<vector<bool>> dpTable(inputLen + 1, vector<bool>(patternLen + 1, false));

    // set the first cell to true: the empty string always matches the empty pattern
    dpTable[0][0] = true;

    // if the input and pattern characters match or the pattern has '?', set the current cell to the value of the cell above and to the left. 
    // the indices correspond to the ith character of the input and jth character of the pattern

    // if the pattern character is '*', set the current cell to the value of the cell above or to the left
    // this is because '*' can match 0 or more characters

    // handle case where input is empty but pattern is not: we set the cell to the value of the cell to the left which, if the pattern is '*', will be true
    // because '*' can match the empty string
    for (int j = 1; j <= patternLen; j++) {
        if (pattern[j - 1] == '*')
            dpTable[0][j] = dpTable[0][j - 1];
    }

    // fill the rest of the table
    for (int i = 1; i <= inputLen; i++) {
        for (int j = 1; j <= patternLen; j++) {
            if (pattern[j - 1] == '?' || pattern[j - 1] == input[i - 1]) // if the current characters match or pattern has '?', set the current cell to the value of the cell above and to the left
                dpTable[i][j] = dpTable[i - 1][j - 1];
            else if (pattern[j - 1] == '*') // if the pattern character is '*', set the current cell to the value of the cell above or to the left
                dpTable[i][j] = dpTable[i - 1][j] || dpTable[i][j - 1]; // the or nicely switches between the two cases of '*' matching 0 or more characters
        }
    }

    // print the dp table
    cout << "DP table: " << endl;
    for (int i = 0; i <= inputLen; i++) {
        for (int j = 0; j <= patternLen; j++) {
            cout << dpTable[i][j] << " ";
        }
        cout << endl;
    }

    return dpTable[inputLen][patternLen];
}

int main() {
    string input = "xzyb";
    string pattern = "x*y";
    int found_pattern_start;
    int found_pattern_end;
    
    long naive_duration = 0;

    bool res = match_wildcard_naive(input, pattern, found_pattern_start, found_pattern_end);
    cout << res << endl;
    for (int i = found_pattern_start; i <= found_pattern_end; i++) {
        cout << input[i];
    }
    cout << endl;

    res = match_wildcard_dp(input, pattern);
    cout << res << endl;




    // int cache_hits = 0;
    // cout << "Running 1000000 tests..." << endl;
    // for (int i = 0; i < 1000000; i++) {
    //     // create a random input
    //     input = "";
    //     for (int j = 0; j < 100; j++) {
    //         // generate random number between 97 and 122 using the system time as a seed
    //         srand(time(NULL) + j);
    //         int num = rand() % 26 + 97;
    //         // convert that number into a character and add it to the input string
    //         input += (char)num;
    //     }

    //     auto start = chrono::high_resolution_clock::now();
    //     match_wildcard_naive(input, pattern, found_pattern_start, found_pattern_end);
    //     auto stop = chrono::high_resolution_clock::now();
    //     auto duration = chrono::duration_cast<chrono::nanoseconds>(stop - start);
    //     naive_duration += duration.count();

    // }
    // cout << naive_duration << endl;
    // // cout << memoized_duration << endl;
    // // average the durations
    // naive_duration /= 1000000;
    // // memoized_duration /= 1000000;

    // cout << "Naive duration: " << naive_duration << " nanoseconds" << endl;
    // cout << "Memoized duration: " << memoized_duration << " nanoseconds" << endl;

    // cout << "Cache hits: " << cache_hits << endl;


    return 0;
}