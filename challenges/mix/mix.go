package mix

// 0

// 0.1 a va b haqiqiy sonlar to'g'ri to'rtburchakning tomonlari bo'lsa uning yuzasi va perimetrini hisoblab ekranga chiqarib beruvchi dastur tuzing.
// P = 2*(a+b)
// S = a*b

// 0.4 Konfetning kilogrami va narxini foydalanuvchidan olib summasini hisoblab bering.

// 0.5 Narxi X so'mdan A kg konfet, narxi Y so'md-an B kg konfetdan qancha summaga farq qilishini topib bering. X,Y,A va B lar foydalanuvchidan qabul qilinsin.

// 0.6 Birinchi avtomobil tezligi v1, ikkinchi avtomobil tezligi v2, ular bir biridan uzoqlashishni boshlasa, t vaqtdan keyin qancha uzoqlikda bo'lishlarini toping.
// S=v*t | S-yo'l, v- tezlik, t-vaqt

// 0.7 Qayiq tezligi v1 bilan berlgilang. Daryo oqimi tezligi v2 bilan. Qayiq oqim bo'yicha t1 soat, oqimga qarshi t2 soat suzgan bo'lsa jami qancha masofa bosganini hisoblovchi dastur tuzing.
// (v1+v2)*t1 + (v1-v2)*t2


// 0.8 Berilgan a va b sonlarining qiymatlarini almashtiring.

// 0.8* Berilgan a va b sonlarining qiymatlarini yangi o'zgaruvchi yaratmasdan almashtiring.

// 0.9 a,b va c sonlari berilgan. a ning qiymatini b ga, b ning qiymatini c ga va c ning qiymatini a ga almashtiring.

// 1 - If

// 1.01 Foydalanuvchi a son kiritsin, siz bu sonni musbat yoki manfiy deb ekranga chiqarib bering.

// 1.02 Foydalanuvchi a son kiritsin, siz bu sonni juft yoki toqligini aniqlab ekranga chiqarib bering.

// 1.03 Uchta son berilgan shulardan kattasini topuvchi dastur tuzing.

// 1.04 Konfet miqdori 10kg dan oshsa qolganini yarim narxida sotiladi. Berilgan narx va og'irligigaga qarab umumiy summani hisoblab bering. Foydalanuvchi tomonidan narx va kg kiritilishi kerak.

// 1.05 Uchta son berilgan shulardan kichkinasini topuvchi dastur tuzing.

// 1.06 Uchta son berilgan shulardan o'rtanchasini topuvchi dastur tuzing.

// 1.07 Berilgan son musbat bo'lsa uni 5 ga ko'paytiring, manfiy bo'lsa 3 ga bo'ling va 0 ga teng bo'lsa 7 ni qo'shing.

// 1. Challenge: Convert a decimal number to binary representation.

// Pseudo Code:
// {
// Input decimalNumber
// Set binary = empty string
// While decimalNumber is not equal to 0
// Set remainder = decimalNumber % 2
// Set binary = remainder + binary
// Set decimalNumber = decimalNumber / 2
// Display "The binary representation is: " + binary
// }

// 2. Challenge: Calculate and display the factorial of a number.

// Pseudo Code:
// Input number
// Set factorial = 1
// For i = 1 to number
// Set factorial = factorial * i
// Display "The factorial of " + number + " is: " + factorial

// 3. Challenge: Check if a string is a palindrome.

// Pseudo Code:
// Input word
// Set isPalindrome = true
// Set length = length of word
// For i = 0 to length/2
// If word[i] != word[length-i-1] Then
// Set isPalindrome = false
// Break
// If isPalindrome Then
// Display word + " is a palindrome"
// Else
// Display word + " is not a palindrome"

// 4. Count the number of vowels in a string.
// Berilgan matndagi unli harflar sonini aniqlovchi funksiya yozing.  aeiou

// Pseudo Code:
// Input word
// Set vowelCount = 0
// For i = 0 to length of word
// If word[i] is a vowel Then
// Increment vowelCount by 1
// Display "The number of vowels in " + word + " is: " + vowelCount

// 5. Reverse a string.

// Pseudo Code:
// Input word
// Set reversedWord = ""
// Set length = length of word
// For i = length-1 to 0
// Append word[i] to reversedWord
// Display "The reversed string is: " + reversedWord

// 6. Check if two strings are anagrams.

// Pseudo Code:
// Input word1, word2
// Sort word1 and store the result in sortedWord1
// Sort word2 and store the result in sortedWord2
// If sortedWord1 is equal to sortedWord2 Then
// Display word1 + " and " + word2 + " are anagrams"
// Else
// Display word1 + " and " + word2 + " are not anagrams"

// 7. Find the largest element in an array.

// Pseudo Code:
// Input array
// Set largest = array[0]
// For i = 1 to length of array
// If array[i] > largest Then
// Set largest = array[i]
// Display "The largest element is: " + largest

// 8. Find the second smallest element in an array.

// Pseudo Code:
// Input array
// Set smallest = array[0]
// Set secondSmallest = array[0]
// For i = 1 to length of array
// If array[i] < smallest Then
// Set secondSmallest = smallest
// Set smallest = array[i]
// Else If array[i] < secondSmallest Then
// Set secondSmallest = array[i]
// Display "The second smallest element is: " + secondSmallest

// 9. Challenge: Calculate the sum of elements in an array.

// Pseudo Code:
// Input array
// Set sum = 0
// For i = 0 to length of array
// Set sum = sum + array[i]
// Display "The sum of elements is: " + sum

// 10. Challenge: Remove duplicates from an array and display the unique elements.

// Pseudo Code:
// Input array
// Set uniqueArray = empty array
// For i = 0 to length of array
// If array[i] is not in uniqueArray Then
// Append array[i] to uniqueArray
// Display "The unique elements are: " + uniqueArray

// 12. Challenge: Sort an array in ascending order.

// Pseudo Code:
// Input array
// Sort array in ascending order
// Display "The sorted array is: " + array

// 13. Challenge: Calculate the average of elements in an array.

// Pseudo Code:
// Input array
// Set sum = 0
// For i = 0 to length of array
// Set sum = sum + array[i]
// Set average = sum / length of array
// Display "The average of elements is: " + average

// 14. Challenge: Calculate the median of elements in an array.

// Pseudo Code:
// Input array
// Sort array in ascending order
// Set middleIndex = length of array / 2
// If length of array is odd Then
// Set median = array[middleIndex]
// Else
// Set median = (array[middleIndex-1] + array[middleIndex]) / 2
// Display "The median of elements is: " + median

// 15. Challenge: Find the index of the first occurrence of an element in an array.

// Pseudo Code:
// Input array, element
// Set index = -1
// For i = 0 to length of array
// If array[i] equals element Then
// Set index = i
// Break
// If index != -1 Then
// Display "The element is found at index: " + index
// Else
// Display "The element is not found in the array"

// 16. Challenge: Check if a number is a perfect square.

// Pseudo Code:
// Input number
// Set squareRoot = square root of number
// If squareRoot is an integer Then
// Display number + " is a perfect square"
// Else
// Display number + " is not a perfect square"

// 17. Challenge: Generate and display all prime numbers up to a given limit.

// Pseudo Code:
// Input limit
// Set primeNumbers = empty array
// For i = 2 to limit
// Set isPrime = true
// For j = 2 to i/2
// If i % j == 0 Then
// Set isPrime = false
// Break
// If isPrime Then
// Append i to primeNumbers
// Display "The prime numbers up to " + limit + " are: " + primeNumbers

// 18. Challenge: Reverse an array in-place (without using additional storage).

// Pseudo Code:
// Input array
// Set start = 0
// Set end = length of array - 1
// While start < end
// Swap array[start] with array[end]
// Increment start by 1
// Decrement end by 1
// Display "The reversed array is: " + array

// 19. Challenge: Find the longest word in a sentence.

// Pseudo Code:
// Input sentence
// Set words = split sentence by spaces
// Set longestWord = words[0]
// For i = 1 to length of words
// If length of words[i] > length of longestWord Then
// Set longestWord = words[i]
// Display "The longest word is: " + longestWord

// 20. Challenge: Count the frequency of each character in a string.

// Pseudo Code:
// Input word
// Set characterCount = empty map/dictionary
// For each character in word
// If character is not in characterCount Then
// Set characterCount[character] = 1
// Else
// Increment characterCount[character] by 1
// Display "Character frequencies:"
// For each key in characterCount
// Display key + ": " + characterCount[key]

// 21. Challenge: Find the common elements between two arrays.

// Pseudo Code:
// Input array1, array2
// Set commonElements = empty array
// For i = 0 to length of array1
// For j = 0 to length of array2
// If array1[i] equals array2[j] Then
// Append array1[i] to commonElements
// Display "The common elements are: " + commonElements

// 22. Challenge: Implement binary search in a sorted array.

// Pseudo Code:
// Input array, element
// Set start = 0
// Set end = length of array - 1
// While start <= end
// Set mid = (start + end) / 2
// If array[mid] equals element Then
// Display "Element found at index: " + mid
// Break
// Else If array[mid] < element Then
// Set start = mid + 1
// Else
// Set end = mid - 1
// If start > end Then
// Display "Element not found in the array"

// 23. Challenge: Check if a string is a pangram (contains all alphabets).

// Pseudo Code:
// Input sentence
// Set alphabets = "abcdefghijklmnopqrstuvwxyz"
// Set lowerCaseSentence = convert sentence to lowercase
// Set isPangram = true
// For each character in alphabets
// If character is not in lowerCaseSentence Then
// Set isPangram = false
// Break
// If isPangram Then
// Display sentence + " is a pangram"
// Else
// Display sentence + " is not a pangram"

// Challenge: Calculate the power of a number using recursion.
// Pseudo Code:

// Input base, exponent
// If exponent == 0 Then
// Return 1
// Else
// Return base * power(base, exponent-1)

// Challenge: Implement selection sort to sort an array in ascending order.

// Pseudo Code:
// Input array
// For i = 0 to length of array - 1
// Set minIndex = i
// For j = i+1 to length of array
// If array[j] < array[minIndex] Then
// Set minIndex = j
// Swap array[i] with array[minIndex]
// Display "The sorted array is: " + array

// Challenge: Find the GCD (Greatest Common Divisor) of two numbers.

// Pseudo Code:
// Input number1, number2
// While number2 is not equal to 0
// Set temp = number2
// Set number2 = number1 % number2
// Set number1 = temp
// Display "The GCD is: " + number1

// Challenge: Check if a number is an Armstrong number.
// Pseudo Code:
// Input number
// Set originalNumber = number
// Set sum = 0
// Set numDigits = number of digits in number
// While number is not equal to 0
// Set digit = number % 10
// Set sum = sum + digit^numDigits
// Set number = number / 10
// If sum equals originalNumber Then
// Display originalNumber + " is an Armstrong number"
// Else
// Display originalNumber + " is not an Armstrong number"

// Challenge: Find the LCM (Least Common Multiple) of two numbers.
// Pseudo Code:
// Input number1, number2
// Set maxNumber = max(number1, number2)
// Set lcm = maxNumber
// While true
// If lcm % number1 == 0 and lcm % number2 == 0 Then
// Display "The LCM is: " + lcm
// Break
// Increment lcm by maxNumber

// Challenge: Implement bubble sort to sort an array in ascending order.
// Pseudo Code:
// Input array
// Set n = length of array
// For i = 0 to n-1
// For j = 0 to n-i-1
// If array[j] > array[j+1] Then
// Swap array[j] with array[j+1]
// Display "The sorted array is: " + array

// Challenge: Check if a string is a valid palindrome ignoring non-alphanumeric characters.
// Pseudo Code:
// Input word
// Set cleanedWord = remove non-alphanumeric characters from word
// Set reversedWord = reverse cleanedWord
// If cleanedWord equals reversedWord Then
// Display word + " is a valid palindrome"
// Else
// Display word + " is not a valid palindrome"

// Challenge: Calculate the sum of digits in a number using recursion.
// Pseudo Code:
// Input number
// If number == 0 Then
// Return 0
// Else
// Return (number % 10) + sumOfDigits(number / 10)

// Challenge: Find the number of occurrences of a substring in a string.
// Pseudo Code:
// Input word, substring
// Set count = 0
// Set wordLength = length of word
// Set subLength = length of substring
// For i = 0 to wordLength - subLength
// If word[i...i+subLength-1] equals substring Then
// Increment count by 1
// Display "The substring occurs " + count + " times"

// Challenge: Convert a decimal number to binary representation.
// Pseudo Code:
// Input decimalNumber
// Set binary = empty string
// While decimalNumber is not equal to 0
// Set remainder = decimalNumber % 2
// Set binary = remainder + binary
// Set decimalNumber = decimalNumber / 2
// Display "The binary representation is: " + binary

// Challenge: Find the length of the longest increasing subarray in an array.
// Pseudo Code:
// Input array
// Set maxLength = 1
// Set length = 1
// For i = 1 to length of array
// If array[i] > array[i-1] Then
// Increment length by 1
// Else
// Set maxLength = max(maxLength, length)
// Set length = 1
// Set maxLength = max(maxLength, length)
// Display "The length of the longest increasing subarray is: " + maxLength

// Challenge: Calculate the square root of a number using the Babylonian method.
// Pseudo Code:
// Input number
// Set guess = number / 2
// Set threshold = 0.0001
// While abs(guess * guess - number) > threshold
// Set guess = (guess + number / guess) / 2
// Display "The square root is: " + guess

// Challenge: Generate and display Pascal's triangle up to a given number of rows.
// Pseudo Code:
// Input numRows
// Set triangle = empty 2D array
// For i = 0 to numRows-1
// Set row = empty array
// For j = 0 to i
// If j equals 0 or j equals i Then
// Set value = 1
// Else
// Set value = triangle[i-1][j-1] + triangle[i-1][j]
// Append value to row
// Append row to triangle
// Display "Pascal's triangle:"
// For each row in triangle
// Display row

// Challenge: Implement merge sort to sort an array in ascending order.
// Pseudo Code:
// Input array
// If length of array <= 1 Then
// Return array
// Set mid = length of array / 2
// Set left = mergeSort(first half of array)
// Set right = mergeSort(second half of array)
// Return merge(left, right)

// Challenge: Find the number of trailing zeros in the factorial of a number.
// Pseudo Code:
// Input number
// Set count = 0
// Set powerOfFive = 5
// While powerOfFive <= number
// Increment count by floor(number / powerOfFive)
// Set powerOfFive = powerOfFive * 5
// Display "The number of trailing zeros is: " + count

// Challenge: Find the first non-repeated character in a string.
// Pseudo Code:
// Input word
// Set characterCount = empty map/dictionary
// For each character in word
// If character is not in characterCount Then
// Set characterCount[character] = 1
// Else
// Increment characterCount[character] by 1
// For each character in word
// If characterCount[character] equals 1 Then
// Display "The first non-repeated character is: " + character
// Break

// Challenge: Check if a number is a palindrome.
// Pseudo Code:
// Input number
// Set originalNumber = number
// Set reversedNumber = 0
// While number is not equal to 0
// Set digit = number % 10
// Set reversedNumber = reversedNumber * 10 + digit
// Set number = number / 10
// If originalNumber equals reversedNumber Then
// Display originalNumber + " is a palindrome"
// Else
// Display originalNumber + " is not a palindrome"

// Challenge: Implement quick sort to sort an array in ascending order.
// Pseudo Code:
// Input array
// If length of array <= 1 Then
// Return array
// Set pivot = last element of array
// Set i = first index of array - 1
// For j = first index of array to last index of array - 1
// If array[j] <= pivot Then
// Increment i by 1
// Swap array[i] with array[j]
// Swap array[i+1] with array[last index of array]
// Set partitionIndex = i + 1
// Set left = quickSort(first part of array)
// Set right = quickSort(second part of array)
// Return concatenate(left, pivot, right)

// Challenge: Calculate the factorial of a number using recursion.
// Pseudo Code:
// Input number
// If number == 0 Then
// Return 1
// Else
// Return number * factorial(number - 1)

// Challenge: Find the number of vowels and consonants in a string.
// Pseudo Code:
// Input word
// Set vowels = "aeiou"
// Set numVowels = 0
// Set numConsonants = 0
// Set lowerCaseWord = convert word to lowercase
// For each character in lowerCaseWord
// If character is in vowels Then
// Increment numVowels by 1
// Else If character is an alphabet Then
// Increment numConsonants by 1
// Display "Number of vowels: " + numVowels
// Display "Number of consonants: " + numConsonants

// Challenge: Find the maximum subarray sum in an array.
// Pseudo Code:
// Input array
// Set maxSum = array[0]
// Set currentSum = array[0]
// For i = 1 to length of array
// Set currentSum = max(array[i], currentSum + array[i])
// Set maxSum = max(maxSum, currentSum)
// Display "The maximum subarray sum is: " + maxSum

// Challenge: Reverse words in a sentence.
// Pseudo Code:
// Input sentence
// Set words = split sentence by spaces
// Set reversedWords = empty array
// For i = length of words - 1 to 0
// Append words[i] to reversedWords
// Set reversedSentence = join reversedWords with spaces
// Display "The reversed sentence is: " + reversedSentence

// Challenge: Implement the Sieve of Eratosthenes to find all prime numbers up to a given limit.
// Pseudo Code:
// Input limit
// Set isPrime = array of size (limit+1), initialized as true
// Set p = 2
// While p^2 <= limit
// If isPrime[p] is true
// For i = p^2 to limit, with a step of p
// Set isPrime[i] = false
// Increment p by 1
// Display "The prime numbers up to " + limit + " are:"
// For i = 2 to limit
// If isPrime[i] is true
// Display i
