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
  , signal : Signal
  , wordlist :  String
  , selectedWord : String
  , synon : Bool
  , showAnswer : Bool
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
  ( Model "" "" Loading "a anywhere below burn climb able apartment bend bus close about appear beneath business clothes above approach beside busy cloud accept area best but coat across arm better buy coffee act around between by cold actually arrive beyond call college add art big calm color admit as bird camera come afraid ask bit can company after asleep bite car completely afternoon at black card computer again attack block care confuse against attention blood careful consider age aunt blow carefully continue ago avoid blue carry control agree away board case conversation ahead baby boat cat cool air back body catch cop alive bad bone cause corner all bag book ceiling count allow ball boot center counter almost bank bore certain country alone bar both certainly couple along barely bother chair course already bathroom bottle chance cover also be bottom change crazy although beach box check create always bear boy cheek creature among beat brain chest cross and beautiful branch child crowd angry because break choice cry animal become breast choose cup another bed breath church cut answer bedroom breathe cigarette dad any beer bridge circle dance anybody before bright city dark anymore begin bring class darkness anyone behind brother clean daughter anything believe brown clear day anyway belong building clearly dead death except funny history law decide excite future hit lay deep expect game hold lead desk explain garden hole leaf despite expression gate home lean die extra gather hope learn different eye gently horse leave dinner face get hospital leg direction fact gift hot less dirt fade girl hotel let disappear fail give hour letter discover fall glance house lie distance familiar glass how life do family go however lift doctor far god huge light dog fast gold human like door father good hundred line doorway fear grab hurry lip down feed grandfather hurt listen dozen feel grandmother husband little drag few grass I local draw field gray ice lock dream fight great idea long dress figure green if look drink fill ground ignore lose drive final group image lot driver finally grow imagine loud drop find guard immediately love dry fine guess important low during finger gun in lucky dust finish guy information lunch each fire hair inside machine ear first half instead main early fish hall interest make earth fit hallway into man easily five hand it manage east fix hang itself many easy flash happen jacket map eat flat happy job mark edge flight hard join marriage eff ort floor hardly joke marry egg flower hate jump matter eight fly have just may either follow he keep maybe else food head key me empty foot hear kick mean end for heart kid meet engine force heat kill member enjoy forehead heavy kind memory enough forest hell kiss mention enter forever hello kitchen message entire forget help knee metal especially form her knife middle even forward here knock might event four herself know mind ever free hey lady mine every fresh hi land minute everybody friend hide language mirror everyone from high large miss everything front hill last moment everywhere full him later money exactly fun himself laugh month moon our quickly send smile more out quiet sense smoke morning outside quietly serious snap most over quite seriously snow mostly own radio serve so mother page rain service soft mountain pain raise set softly mouth paint rather settle soldier move pair reach seven somebody movie pale read several somehow much palm ready sex someone music pants real shadow something must paper realize shake sometimes my parent really shape somewhere myself part reason share son name party receive sharp song narrow pass recognize she soon near past red sheet sorry nearly path refuse ship sort neck pause remain shirt soul need pay remember shoe sound neighbor people remind shoot south never perfect remove shop space new perhaps repeat short speak news personal reply should special next phone rest shoulder spend nice photo return shout spin night pick reveal shove spirit no picture rich show spot nobody piece ride shower spread nod pile right shrug spring noise pink ring shut stage none place rise sick stair nor plan river side stand normal plastic road sigh star north plate rock sight stare nose play roll sign start not please roof silence state note pocket room silent station nothing point round silver stay notice police row simple steal now pool rub simply step number poor run since stick nurse pop rush sing still of porch sad single stomach off position safe sir stone offer possible same sister stop office pour sand sit store officer power save situation storm often prepare say six story oh press scared size straight okay pretend scene skin strange old pretty school sky street on probably scream slam stretch once problem screen sleep strike one promise sea slide strong only prove search slightly student onto pull seat slip study open push second slow stuff or put see slowly stupid order question seem small such other quick sell smell suddenly suggest thick tree wash window suit thin trip watch wine summer thing trouble water wing sun think truck wave winter suppose third true way wipe sure thirty trust we wish surface this truth wear with surprise those try wedding within sweet though turn week without swing three twenty weight woman system throat twice well wonder table through two west wood take throw uncle wet wooden talk tie under what word tall time understand whatever work tea tiny unless wheel world teach tire until when worry teacher to up where would team today upon whether wrap tear together use which write television tomorrow usual while wrong tell tone usually whisper yard ten tongue very white yeah terrible tonight view who year than too village whole yell thank tooth visit whom yellow that top voice whose yes the toss wait why yet their touch wake wide you them toward walk wife young themselves town wall wild your then track want will yourself there train war win these travel warm wind" "" False False
  , Random.generate Initgame (Random.int 0 999)
    
  )




----UPDATE----


type Msg
  = Initgame Int
  | Chooseword
  | GotQuote (Result Http.Error (List (List Meaning)))
  | Answer String
  | ShowWord
  | ShowSynon




update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of 
    Initgame index ->
      ( infor_word model index
      )
    
    Chooseword ->
      ({model | showAnswer = False}, Random.generate Initgame (Random.int 0 999))
     
    GotQuote result ->
      case result of
        Ok quote ->
          ({model | signal = Success quote}, Cmd.none)
        Err err ->
          ({model | signal = Failure err}, Cmd.none)
          
    Answer usranswer ->
      ({model | answer=usranswer}, Cmd.none)
      
    ShowWord ->
      ({model | showAnswer = not model.showAnswer}, Cmd.none)
    
    ShowSynon ->
      ({model | synon = not model.synon}, Cmd.none)


    
infor_word : Model -> Int -> (Model, Cmd Msg)
infor_word model index =
  let
    url = "https://api.dictionaryapi.dev/api/v2/entries/en/" ++ get_word model.wordlist index
    request = Http.get
      { url = url
      , expect = Http.expectJson GotQuote quoteDecoder
      }
  in
  ( { model | answer = "", synon = False, selectedWord = get_word model.wordlist index }
  , request
  )

    
get_word : String -> Int -> String
get_word words nbr = 
  let 
    words_list = String.words(words)
    wordarray = Array.fromList words_list
    word = Array.get nbr wordarray 

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
      , viewSelectedWord model.showAnswer model.selectedWord
      , ol[][ h3[][text "meanings"], getDefinition quote 0 model.synon]
      , viewValidate model
      , viewInput "text" "Answer" model.answer Answer
      , checkbox ShowWord "show the answer" model.showAnswer
      , div[style "padding-left" "80px"]  
        [ button [ onClick Chooseword, style "display" "block" ] [text "More game"]]
      , div[style "padding-left" "80px"]  
        [ button [ onClick ShowSynon, style "display" "block" ] [text "Synonyms"]]
      ]

viewSelectedWord : Bool-> String -> Html Msg
viewSelectedWord showanswer selectedWord =
  case showanswer == True of
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
    

checkbox : msg -> String -> Bool -> Html msg
checkbox msg name isChecked =
  label
    [ style "padding-left" "100px" ]
    [ input [ type_ "checkbox", checked isChecked, onClick msg ] []
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
    pre[][text("Synonyms : "), text (String.join ", " (meaning.synonyms) )]
  else
    pre[][]
















    
    
    
    
    
    
    
    
    
    

    
    
    
    
    
    
    
    
    
    