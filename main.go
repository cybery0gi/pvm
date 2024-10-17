package main

import (
    "fmt"
)

// Map of Unicode values (without the \u0 prefix) to Sanskrit Devanagari characters
var unicodeToDevanagari = map[string]string{
    "905": "अ",
    "906": "आ",
    "907": "इ",
    "908": "ई",
    "909": "उ",
    "90A": "ऊ",
    "90B": "ऋ",
    "960": "ॠ",
    "90C": "ऌ",
    "961": "ॡ",
    "90F": "ए",
    "910": "ऐ",
    "913": "ओ",
    "914": "औ",
    "90502": "अं",
    "90503": "अः",
    "915": "क",
    "916": "ख",
    "917": "ग",
    "918": "घ",
    "919": "ङ",
    "91A": "च",
    "91B": "छ",
    "91C": "ज",
    "91D": "झ",
    "91E": "ञ",
    "91F": "ट",
    "920": "ठ",
    "921": "ड",
    "922": "ढ",
    "923": "ण",
    "924": "त",
    "925": "थ",
    "926": "द",
    "927": "ध",
    "928": "न",
    "92A": "प",
    "92B": "फ",
    "92C": "ब",
    "92D": "भ",
    "92E": "म",
    "92F": "य",
    "930": "र",
    "932": "ल",
    "935": "व",
    "936": "श",
    "937": "ष",
    "938": "स",
    "939": "ह",
    "915094D937": "क्ष",
    "924094D930": "त्र",
    "91C094D91E": "ज्ञ",
}

func main() {
    // Test printing Unicode values and their corresponding Devanagari characters
    for unicodeValue, character := range unicodeToDevanagari {
        fmt.Printf("Unicode: %s, Character: %s\n", unicodeValue, character)
    }
}
