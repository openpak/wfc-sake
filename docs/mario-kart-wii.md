# Mario Kart Wii on wfc-sake

What the server answers (game id 1687, secret in `game_list.tsv`):

- Sake tables `mariokartwii/FriendInfo` (per-player info blob, Mii data filtered for non-owners),
  `GhostData` (the nearest slower ghost: filter `course = N and gameid = 1687 and time < T`) and
  `StoredGhostData` (the best ghost of a course and region).
- File server: `download.aspx` (by `fileid`), `ghostdownload.aspx` (`region=0&p0=<pid>&c0=<course>&t0=<time>`,
  the nearest slower ghost of another player, with the 0x200 byte download header) and the ghost upload
  (`regionid, courseid, score, pid, playerinfo, contest`).
- Rankings are served by wfc-race (`GetTopTenRankings`), from the table this service writes.

Since wfc-common v0.2.0 every accepted upload with a body is also kept in `mario_kart_wii_ghosts`
with its RKG header metadata, for the wfc-race read API. The game never reads that table.

## Parameters the server logs but does not know

Any query parameter on `ghostdownload.aspx` or the ghost upload beyond the ones above is logged at
info level by name and length only ("Unknown query parameter"). Nothing in the public docs
(Wiimm, wiki.tockdom.com, WiiLink) describes a ghost download by rank or by friend, or competition
ghost bodies, so none is implemented; the server already ranks competition ids above 31 (uploads with
`contest`, stored without a body).

## Console tests needed

1. Time Trial on any course, finish, upload: confirm the ghost appears in wfc-race
   `GET /api/mkw/v1/courses/<n>/ghosts` with the right vehicle, character, date and lap times.
2. Mario Kart Channel, Worldwide ranking, select a row and download its ghost: check wfc-sake log for
   `Unknown query parameter` or an unhandled table/filter and report the names.
3. Mario Kart Channel, competition list and a competition ghost: same log check; any unknown table or
   parameter is the missing piece.
