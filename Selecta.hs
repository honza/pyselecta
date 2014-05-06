{-# LANGUAGE OverloadedStrings #-}

import qualified Data.ByteString as BS
import qualified Data.ByteString.Char8 as B
import System.Posix.Env.ByteString
import Data.Word
import Data.Maybe
import Data.Char
import Data.List
import Data.Ord

stringIndex :: BS.ByteString -> Word8 -> Int -> Maybe Int
stringIndex s char offset = fmap (+ offset) $ BS.elemIndex char (BS.drop offset s)

findEndOfMatch :: BS.ByteString -> [Word8] -> Int -> Maybe Int
findEndOfMatch s [] findex = Just findex
findEndOfMatch s [c] findex = stringIndex s c (findex + 1)
findEndOfMatch s (c:cs) findex =
    case r of
        Nothing -> Nothing
        (Just ints) -> Just (last ints)
        where r = sequence $ result:[findEndOfMatch s cs newIndex]
                where result = findEndOfMatch s [c] findex
                      newIndex = findex + 1

findCharInString :: BS.ByteString -> Word8 -> [Int]
findCharInString s c = map pred (BS.elemIndices c s)

computeMatchResult :: Maybe Int -> Int -> Maybe Int
computeMatchResult Nothing _ = Nothing
computeMatchResult (Just last_index) first_index =
        Just (last_index - first_index + 1)

computeMatchLength :: BS.ByteString -> [Word8] -> Int
computeMatchLength s (c:cs) =
    case m of
        Nothing -> 0
        Just v -> case v of
                      [] -> 0
                      x -> minimum x
    where m = sequence $ filter isJust result
          result =
            map (\f ->
                    computeMatchResult (findEndOfMatch s cs f) f)
                 (findCharInString s c)

bsScore :: BS.ByteString -> BS.ByteString -> Double
bsScore choice query =
    if match_length == 0
        then 0.0
        else score / fromIntegral (BS.length choice)

    where match_length = computeMatchLength choice (BS.unpack query)
          score = fromIntegral (BS.length query) / fromIntegral match_length

lower :: BS.ByteString -> BS.ByteString
lower = B.map toLower

score :: [Word8] -> [Word8] -> Double
score _ [] = 1.0
score [] _ = 0.0
score a b = bsScore (packLow a) (packLow b)
    where packLow = lower . BS.pack

match :: [BS.ByteString] -> BS.ByteString -> [BS.ByteString]
match files query = reverse $ map fst (sortBy (comparing snd) (filter ((<) 0 . snd) scores))
    where scores = map (\s -> (s, score (BS.unpack s) (BS.unpack query))) files

main = do
    args <- getArgs

    let query = case args of
            [] -> error "Need at least of argument"
            [x] -> x
            (x:xs) -> x

    files <- BS.getContents
    B.putStrLn $ B.unlines $ match (B.lines files) query
