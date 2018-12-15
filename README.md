# Trumango

Trumango is a basic question answering system based on pattern matching. This project written in Golang with TDD style.

# Usage

Trumango gets question sentences from file which has one question sentence per line and gets text which will be used for searching answers, from another file. Text must be written in one line.

#### Example Question File

```
    With an apathetic shrug, what does Truman replace?
    He picks up the framed picture of his wife from where?
    The sound of the children triggers what in his head?
    ...
```

Questions file set to `questions.txt` and text file set to `the_truman_show_script.txt` by default.

Question and text files can be specified using flags.

Show usage and flag descriptions: `./trumango --help`

##### Specifying custom question and text file:

```
    ./trumango -q my_questions.txt -t my_text.txt
```

# How it works?

## NLP Module

`nlp` module contains following functionalities for text processing.

- `Stem` wraps stemming functionality of [porter2](https://github.com/dchest/stemmer/porter2) stemmer. `Stem` finds the stems of each word using porter2 stemmer then assembles the sentence again.
- `ClearStopWords` uses the [stopwords](https://github.com/bbalet/stopwords) library for clearing stop words.
- `SplitSentences` uses the segmentation functionality of [prose](https://github.com/jdkato/prose) to split a text into sentences accurately.

## Horspool Module

`horspool` module contains following functionalities for pattern matching with Horspool Algorithm which has **O(n)** time complexity in average.

- `Find` finds the index of first matching pattern in text using horspool algorithm.
- `FindLast` finds the index of last matching pattern in text using horspool algorithm.

## Util Module

`util` module contains `Difference` function which finds the difference of given string array A from given string array B.

---

## Getting and Processing Input

Trumango loads specified files, then places each question into a map -_which will be filled with corresponding answer sentences of questions_- and splits the text into sentences by using [prose](https://github.com/jdkato/prose). Prose has very extensive and advanced features such as tokenizing a sentence and tagging each word as verb, noun etc. but in this project prose just used for splitting the text into sentences more accurately.

## Finding Answer Sentences of Questions

After the processing of input, for each question in the map, corresponding question's stop words cleared and stemmed for searching in the sentences. When clearing and stemming process of question is done, question splitted into words and each word searcher in cleared and stemmed sentence for match. If match percentage of question words is higher then desired percentage -_which can be specified by `-m` flag_-, then that sentence assigned to map for corresponding question.

## Finding Exact Answers

When finding answer sentence process is done, by using question and answer sentence map, exact answers will be found. For this process, sentence is cleared and stemmed but for retrieving the original version of word from stemmed word, a map of stemmed words and original words constructed while stemming operation. After this process, by getting the difference of cleared and stemmed sentence words from cleared and stemmed question words, we end up with exact answer words.

## Example

#### Original Question:

`The sound of the children triggers what in his head?`

#### Stop Words Filtered and Stemmed Question:

`sound children trigger head`

#### Original Answer Sentence:

`The sound of the children triggers a memory in his head.`

#### Stop Words Filtered and Stemmed Answer Sentence:

`sound children trigger memori head`

#### Difference:

`memori`

#### Original word:

`memory`

### Further examples can be found in `truman_test.go` file.

---

### Full list of libraries used:

- https://github.com/bbalet/stopwords
- https://github.com/dchest/stemmer/porter2
- https://github.com/jdkato/prose
- https://github.com/fatih/color (for coloring output)
