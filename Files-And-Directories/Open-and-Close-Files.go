package main

import (
   "log"
   "os"
)

func main() {
   // Simple read only open. We will cover actually reading
   // and writing to files in examples further down the page
   file, err := os.Open("test.txt")
   if err != nil {
      log.Fatal(err)
   }
   file.Close()

   // OpenFile with more options. Last param is the permission mode
   // Second param is the attributes when opening
   file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
   if err != nil {
      log.Fatal(err)
   }
   file.Close()

   // Use these attributes individually or combined
   // with an OR for second arg of OpenFile()
   // e.g. os.O_CREATE|os.O_APPEND
   // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

   // os.O_RDONLY // Read only
   // os.O_WRONLY // Write only
   // os.O_RDWR // Read and write
   // os.O_APPEND // Append to end of file
   // os.O_CREATE // Create is none exist
   // os.O_TRUNC // Truncate file when opening
}


References:
https://golang.org/pkg/os/#Open
https://golang.org/pkg/os/#File.Close

Questions answered by this page:How do you open a file for reading in go lang?
How do you close a file in go lang?



<ins class="adsbygoogle"
     style="display:block"
     data-ad-client="ca-pub-4058794840952844"
     data-ad-slot="4835580402"
     data-ad-format="auto">

(adsbygoogle = window.adsbygoogle || []).push({});



   
 


  
    Related ExamplesBuffered ReaderBuffered WriterChange File PermissionsChange The Current DirectoryCheck If File ExistsCompress a FileCopy A FileCreate A Temporary DirectoryCreate A Temporary FileCreate Empty FileCreate a Zip ArchiveCtime Mtime AtimeDelete A FileExtract Zip ArchiveGet File InfoGet File Mode And Permission BitsGet The Current Working DirectoryHashing and ChecksumsOpen and Close FilesRead All Bytes in a FileRead File ContentsRename A FileSeek Positions in FileTruncate A FileUncompress a FileWrite Bytes to a FileWrite to a file


<ins class="adsbygoogle"
     style="display:block"
     data-ad-client="ca-pub-4058794840952844"
     data-ad-slot="1144086401"
     data-ad-format="auto">

(adsbygoogle = window.adsbygoogle || []).push({});


27
   
 

 


  
  







