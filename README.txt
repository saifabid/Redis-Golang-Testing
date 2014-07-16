Supports del,set,incr,get,expire,setex,ttl

 Could've made it alot easier and just let redis handle errors and responses.
 However I wanted to practice parsing strings etc etc and so 
did manual checks on top of the default redis ones 
(hence why only limited commands are accpeted for now)