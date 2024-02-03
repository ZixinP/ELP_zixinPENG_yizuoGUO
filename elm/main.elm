module Main exposing (Model, Msg(..), init, main, update, view)

import Browser exposing (..)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Random
import Array exposing (..)
import Dict exposing (Dict)
import Json.Decode exposing (Decoder, map4, map3, field, int, string, list, decodeString)



-----MAIN----

main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-----MODEL-----


type alias Model = 
  { text : String
  , answer : String
  , result : Bool
  , signal : Signal
  , wordlist : List String
  , selectedWord : String
  , synon : Bool
  }
  
type Signal
  = Failure Http.Error
  | Loading
  | Success (List (List Meaning))

type alias Meaning = 
  { partOfSpeech : String
  , definitions : (List String)
  , synonyms : (List String)
  }

init : () -> (Model, Cmd Msg)
init _ =
  ( Model "" "" False Loading [] "" False
  , Http.get{ 
    url="https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt"
    , expect = Http.expectString GotVocabularies 
    }
 
  )




----UPDATE----


type Msg
  = Initgame Int
  | Chooseword
  | GotQuote (Result Http.Error (List (List Meaning)))
  | Answer String
  | ShowWord
  | GotVocabularies (Result Http.Error String)
  | ShowSynon




update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotVocabularies result ->
      case result of
        Ok vocabularies ->
          let
            wordlist = String.words vocabularies
          in
            ({ model | wordlist = wordlist }, Cmd.none)
        Err err ->
          ({model | signal = Failure err},Cmd.none)
      
    Initgame index ->
      ( {model | selectedWord = get_word model.wordlist index}, 
        infor_word model index
      )
    
    Chooseword ->
      (model, Random.generate Initgame (Random.int 0 999))
     
    GotQuote result ->
      case result of
        Ok quote ->
          ({model | signal = Success quote}, Cmd.none)
        Err err ->
          ({model | signal = Failure err}, Cmd.none)
          
    Answer usranswer ->
      ({model | answer=usranswer}, Cmd.none)
      
    ShowWord ->
      ({model | result = not model.result}, Cmd.none)
    
    ShowSynon ->
      ({model | synon = not model.synon}, Cmd.none)


    
infor_word : Model -> Int -> Cmd Msg
infor_word model int =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ (get_word model.wordlist int)
    , expect = Http.expectJson GotQuote quoteDecoder
    }
    
get_word : List String -> Int -> String
get_word words nbr = 
  let
    word = getAt nbr words
    
    -- wordarray = Array.fromList words --
    -- word = Array.get nbr wordarray --

  in
    Maybe.withDefault "no word" word
  
  
    
 -- Decode Json
quoteDecoder : Decoder (List (List Meaning))
quoteDecoder =
    (Json.Decode.list typeMeaningsDecoder) 

typeMeaningsDecoder : Decoder (List Meaning)
typeMeaningsDecoder = 
    (field "meanings" listMeaningDecoder)

listMeaningDecoder : Decoder (List Meaning)
listMeaningDecoder = 
    Json.Decode.list meaningDecoder

meaningDecoder : Decoder Meaning
meaningDecoder = 
  map3 Meaning
    (field "partOfSpeech" string)
    (field "definitions" (Json.Decode.list definitionDecoder))
    (field "synonyms" (Json.Decode.list string))
   
definitionDecoder : Decoder String
definitionDecoder = 
    (field "definition" string)








-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  case model.signal of
    -- to see what reason cause the failure
    Failure err ->
      case err of
        Http.BadUrl string ->
          div []
            [ text "Error in loading the quote."
            , pre[][text "\n"]
            , text ("BadUrl: "++string)
            , button [ onClick Chooseword, style "display" "block" ] [text "try again"]
            ]
        Http.Timeout ->
          div []
            [ text "Error in loading the quote."
            , pre[][text "\n"]
            , text ("Timeout")
            , button [ onClick Chooseword, style "display" "block" ] [text "try again"]
            ]
        Http.NetworkError ->
          div []
            [ text "Error in loading the quote."
            , pre[][text "\n"]
            , text ("NetworkError")
            , button [ onClick Chooseword, style "display" "block" ] [text "try again"]
            ]
        Http.BadStatus int ->
          div []
            [ text "Error in loading the quote."
            , pre[][text "\n"]
            , text ("BadStatus: "++(String.fromInt int))
            , button [ onClick Chooseword, style "display" "block" ] [text "try again"]
            ]
        Http.BadBody string ->
          div []
            [ text "Error in loading the quote."
            , pre[][text "\n"]
            , text ("BadBody: "++string)
            , button [ onClick Chooseword, style "display" "block" ] [text "try again"]
            ]
            
    Loading ->
      text "Loading..."

    Success quote->
      div[style "padding-left" "200px"]
      [ viewTitle model
      , viewSelectedWord model.result model.selectedWord
      , ol[][ h3[][text "meanings"], getDefinition quote 0 model.synon]
      , viewValidate model
      , viewInput "text" "Answer" model.answer Answer
      , checkbox ShowWord "show the answer"
      , div[style "padding-left" "80px"] 
        [ button [ onClick Chooseword, style "display" "block"] [text "Start game"]]
      , div[style "padding-left" "80px"]  
        [ button [ onClick Chooseword, style "display" "block" ] [text "More game"]]
      , div[style "padding-left" "80px"]  
        [ button [ onClick ShowSynon, style "display" "block" ] [text "Synonyms"]]
      ]

viewSelectedWord : Bool-> String -> Html Msg
viewSelectedWord result selectedWord =
  case result == True of
    True ->
      div[style "padding-left" "80px"] [ text selectedWord ]
    False ->
      div[] []


viewTitle : Model -> Html msg
viewTitle model = 
    h1 [] [ text "Guess the word" ]
    
   
viewInput : String -> String -> String -> (String -> msg) -> Html msg
viewInput text p input_word toMsg =
  div[style "padding-left" "80px"] [ input [ type_ text, placeholder p, value input_word, onInput toMsg ] [] ]
  
  
viewValidate : Model -> Html msg
viewValidate model =
  if String.toLower(model.answer) == model.selectedWord then
    div [ style "padding-left" "80px", style "color" "green" ] [ text "You find it !!! " ]
  else
    div [ style "padding-left" "80px", style "color" "red" ] [ text "wrong word ! " ]
    

checkbox : msg -> String -> Html msg
checkbox msg name =
  label
    [ style "padding-left" "100px" ]
    [ input [ type_ "checkbox", onClick msg ] []
    , text name
    ]
    

    


-- get partOfSpeechs , definisions and synonyms 


getDefinition : (List (List Meaning)) -> Int -> Bool -> Html Msg
getDefinition quote int showsyn = 
  let
    array_ListMeaning = fromList(quote)            --list of all meanings
    maybe_ListMeaning = (get 0 array_ListMeaning)
    listMeaning = Maybe.withDefault [] maybe_ListMeaning
    array_Meaning = fromList(listMeaning)          -- list of partOfSpeech, definitions and synonyms
    maybe_Meaning = (get int array_Meaning)
    meaning = Maybe.withDefault (Meaning "" [""][""]) maybe_Meaning
    array_Definition = fromList(meaning.definitions)
  in
    if meaning /= Meaning "" [""] [""]then
      div[]
      [ ol[][ text meaning.partOfSpeech
            , ol[][ getAllDefinition array_Definition 0 ]
            , ol[][ showSynon showsyn meaning]
            ]
      , getDefinition quote (int+1) showsyn
      ]
    else
      div[][]

getAllDefinition : Array String -> Int -> Html Msg
getAllDefinition array_Definition nbr = 
  let
    maybe_Definition = (get nbr array_Definition)
    definition = Maybe.withDefault "" maybe_Definition
  in
    if definition /= "" then
      pre[] [text (String.fromInt (nbr+1) ++ ". ")
      , text definition
      , getAllDefinition array_Definition (nbr+1) 
      ]
    else 
      pre[] [text (Maybe.withDefault "" maybe_Definition)]

showSynon : Bool -> Meaning -> Html Msg
showSynon showsyn meaning = 
  if showsyn /= False then
    pre[][text("Synonyms"), text (String.join ", " (meaning.synonyms) )]
  else
    pre[][]
    
getAt : Int -> List a -> Maybe a
getAt index list =
    case List.drop index list of
        [] ->
            Nothing

        (x :: _) ->
            Just x    
















    
    
    
    
    
    
    
    
    
    

    
    
    
    
    
    
    
    
    
    