# advent2023

Using Go cuz I've never used it before!

**Day 1**

    P1 wasn't anything interesting. Cheesed out and used regex, but don't think it was completely out of place here.

    P2 I reached for a high-low pointer solve, not sure if this is the kinda thing thats "idiomatic" Go.  Don't really like my usage of map either (slice was more straightforward), but at quick glace didn't grok what Go is doing when I couldn't find a vector, so used what worked.

**Day 2**

    P1 came simple enough. Retained the full concepts of games/hands/counts in case it gained relevance by P2, but should have just broke the game down into counts. Looking back its also a bit odd that I chose to parse colors by length instead of text, but only just looks funky, no big problems

    P2 basically had the same features of P1, just reused strats.  Same issues, but still worked well.

**Day 3**

    P1 I saw this problem and immediately started thinking back to tactics used when dealing with linear camera image processing for automation, where the goal was to stuff work between each frame instead of storing all 3k+ lines of it at once and hoping it could be processed before action needed to be taken.  Also took pains to make sure numbers only counted once in the instance they were in range of more than one symbol (doesn't look like it ever happened though).

    P2 exact same as P1, just invert outer loop so "gears" that touch the same number WOULD count them twice in case such a situation ever came up (never did)

**Day 4**

    P1 had a silly impulse to parse my own digits instead of using normal language tools.  Also brute force tested tries vs wins, but data didn't look like any set was going to be large enough for this to be an issue.

    P2 created a rolling multiplier list, consume front to determine how many additional copies of current card, add a new 0 to the end.  Like how that turned out.

**Day 5**

    P1 feel like I botched this one in a way. Didn't pay close enough attention to the data, so treated it like "line by line" instead of breaking into obvious "\n\n" chunks. Lead to the weird parsing steps and states.  After that, stored each section in map to start name just in case P2 ended up doing something with ending each seed on a different category than "location", but instead should have made more strict ordering.

    P2 it seemed obvious to me that re-writing the transform traversal to return multiple "start/len" pairs to minimize how many full traversals would occur, but I was getting pretty tired and thought I would at least check what brute force looked like.  3 minutes later I had my answer.  I'm not going to do these at midnight anymore and will wait till I wake up, I really don't like the cheap answer here.

**Day 6**

    P1 formula came really quickly, just a bit of fudging around when the parabola intersects are integer values themselves (saw later some very good solutions, better than mine for handling this situation).  Still, came nice and easy.  Also didn't feel like getting further into string parsing tools, and given the inconsistent positioning of items in the input, ran my own integer parser.

    P2 since handling via simple math algorithm, just had to do an even simpler input parse and run the same logic, came together in no time.
