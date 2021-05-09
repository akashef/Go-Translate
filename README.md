# Go-Translate
Free Unlimited Google Translate API

Case 1 :
 Send Objects of Array Strings
 <h2>CASE 1:</h2>
  <h4>Input:</h4>
  Text = {"EN": ["Try","Translate","Words"]}<br>
  Langs = {"AR", "FR", "SP"}
  <h2>Output</h2>
     { "AR": ["جرب", "ترجمة" ,"الكلمات"] ,<br> 
       "EN":[ "Try", "Translate" ,"Words"], <br>
       "FR":["Essayez" de traduire des mots],<br>
       "SP":["Try", "Translate" ,"Words"]<br>
      }<br>
 <h2>CASE 2:</h2>
  <h4>Input:</h4>
  Text = {"EN": "Try Translate Words"}<br>
  Langs = {"AR", "FR", "SP"}
  <h2>Output</h2>
     { "AR":"جرب ترجمة الكلمات",<br> 
       "EN":"Try Translate Words", <br>
       "FR":"Essayez de traduire des mots,<br>
       "SP":"Try Translate Words"<br>
      }<br>
       <h2>CASE 3:</h2>
  <h4>Input:</h4>
  Text =  "Try Translate Words" <br>
  InputLangs = EN
  OutputLangs= FR
  <h2>Output</h2>
    Essayez de traduire des mots
 
