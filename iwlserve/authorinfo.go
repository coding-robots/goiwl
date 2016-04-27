package main

// TODO(dchest): put this stuff in templates or JSON instead.
import "html/template"

type AuthorInfo struct {
	Name             string
	PictureFile      string
	PictureCopyright template.HTML
	WikiLink         string
	Desc             template.HTML
}

var authorInfos = []*AuthorInfo{
	{
		"Agatha Christie",
		"agatha_christie.jpg",
		"", // Public domain, plaque
		"https://en.wikipedia.org/wiki/Agatha_Christie",
		`Dame Agatha Mary Clarissa Christie DBE (née Miller; 15
September 1890 – 12 January 1976) was a British crime writer of novels, short
stories, and plays. She also wrote romances under the name Mary Westmacott, but
she is best remembered for her 66 detective novels and more than 15 short story
collections (especially those featuring Hercule Poirot or Miss Jane Marple), and
her successful West End plays.
<p>
According to the <cite>Guinness Book of World Records</cite>, Christie is the
best-selling novelist of all time. Her novels have sold roughly four billion
copies, and her estate claims that her works rank third, after those of William
Shakespeare and the Bible, as the most widely published books. According to
<cite>Index Translationum</cite>, Christie is the most translated individual
author, with only the collective corporate works of Walt Disney Productions
surpassing her. Her books have been translated into at least 103 languages.
<p>
Agatha Christie published two autobiographies: a posthumous one covering
childhood to old age; and another chronicling several seasons of archaeological
excavation in Syria and Iraq with her second husband, archaeologist Max
Mallowan. The latter was published in 1946 with the title, <cite>Come, Tell Me
How You Live.</cite>
<p>
Christie's stage play <cite>The Mousetrap</cite> holds the record for the
longest initial run: it opened at the Ambassadors Theatre in London on 25
November 1952 and as of 2012 is still running after more than 24,600
performances. In 1955, Christie was the first recipient of the Mystery Writers
of America's highest honour, the Grand Master Award, and in the same year
Witness for the Prosecution was given an Edgar Award by the MWA for Best Play.
Many of her books and short stories have been filmed, some more than once
(<cite>Murder on the Orient Express</cite>, <cite>Death on the Nile</cite> and
<cite>4.50 from Paddington</cite> for instance), and many have been adapted for
television, radio, video games and comics.`,
	},
	{
		"Anne Rice",
		"anne_rice.jpg",
		"", // Public domain, source: http://en.wikipedia.org/wiki/File:Anne_Rice.jpg
		"https://en.wikipedia.org/wiki/Anne_Rice",
		`Anne Rice (born Howard Allen Frances O'Brien; October 4, 1941) is a best-selling American author of metaphysical gothic fiction, Christian literature and erotica from New Orleans, Louisiana. Her books have sold nearly 100 million copies, making her one of the most widely read authors in modern history.

		<p>Bibliography:
		<ul>
		<li>Angel Time (2009)
		<li>Blackwood Farm (2002)
		<li>Blackwood Farm (2002)
		<li>Blood Canticle (2003)
		<li>Blood Canticle (2003)
		<li>Blood and Gold (2001)
		<li>Christ the Lord: Out of Egypt (2005)
		<li>Christ the Lord: The Road to Cana (2008)
		<li>Cry to Heaven (1982)
		<li>Interview with the Vampire (1976)
		<li>Lasher (1993)
		<li>Memnoch the Devil (1995)
		<li>Merrick (2000)
		<li>Merrick (2000)
		<li>Of Love and Evil (2010)
		<li>Pandora (1998)
		<li>Servant of the Bones (1996)
		<li>Taltos (1994)
		<li>The Feast of All Saints (1979)
		<li>The Mummy (1989)
		<li>The Queen of the Damned (1988)
		<li>The Tale of the Body Thief (1992)
		<li>The Vampire Armand (1998)
		<li>The Vampire Lestat (1985)
		<li>The Witching Hour (1990)
		<li>The Wolf Gift (2012)
		<li>Violin (1997)
		<li>Vittorio the Vampire (1999)
		</ul>
		`,
	},
	{
		"Arthur Clarke",
		"arthur_clarke.jpg",
		"Photo by Amy Marash. Public domain.", // Public domain
		"https://en.wikipedia.org/wiki/Arthur_C._Clarke",
		`Sir Arthur Charles Clarke, CBE, FRAS, Sri Lankabhimanya, (16 December 1917 – 19 March 2008) was a British science fiction author, inventor, and futurist, famous for his short stories and novels, among them 2001: A Space Odyssey (1968), and as a host and commentator in the British television series Mysterious World. For many years, Robert A. Heinlein, Isaac Asimov, and Clarke were known as the "Big Three" of science fiction.
<p>
Clarke served in the Royal Air Force as a radar instructor and technician from 1941–1946. He proposed a satellite communication system in 1945 which won him the Franklin Institute Stuart Ballantine Gold Medal in 1963. He was the chairman of the British Interplanetary Society from 1947–1950 and again in 1953.
<p>
Clarke emigrated to Sri Lanka in 1956 largely to pursue his interest in scuba diving; that year, he discovered the underwater ruins of the ancient Koneswaram temple in Trincomalee. He lived in Sri Lanka until his death. He was knighted by Queen Elizabeth II in 1998, and was awarded Sri Lanka's highest civil honour, Sri Lankabhimanya, in 2005.`,
	},
	{
		"Arthur Conan Doyle", //TODO add more text to bio.
		"arthur_conan_doyle.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Arthur_Conan_Doyle",
		`Sir Arthur Ignatius Conan Doyle DL (22 May 1859 – 7 July 1930) was a Scottish physician and writer, most noted for his stories about the detective Sherlock Holmes, generally considered a milestone in the field of crime fiction, and for the adventures of Professor Challenger. He was a prolific writer whose other works include science fiction stories, plays, romances, poetry, non-fiction and historical novels.`,
	},
	{
		"Bram Stoker", //TODO add more text to bio.
		"bram_stoker.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Bram_Stoker",
		`Abraham "Bram" Stoker (8 November 1847 – 20 April 1912) was an Irish novelist and short story writer, best known today for his 1897 Gothic novel Dracula. During his lifetime, he was better known as personal assistant of actor Henry Irving and business manager of the Lyceum Theatre in London, which Irving owned.`,
	},
	{
		"Charles Dickens",
		"charles_dickens.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Charles_Dickens",
		`Charles John Huffam Dickens (7 February 1812 – 9 June 1870) was an English writer and social critic who is generally regarded as the greatest novelist of the Victorian period and the creator of some of the world's most memorable fictional characters. During his lifetime Dickens' works enjoyed unprecedented popularity and fame, but it was in the twentieth century that his literary genius was fully recognized by critics and scholars. His novels and short stories continue to enjoy an enduring popularity among the general reading public.
<p>
Dickens rocketed to fame with the 1836 serial publication of The Pickwick Papers. Within a few years he had become an international literary celebrity, celebrated for his humour, satire, and keen observation of character and society. His novels, most published in monthly or weekly instalments, pioneered the serial publication of narrative fiction, which became the dominant Victorian mode for novel publication. The instalment format allowed Dickens to evaluate his audience's reaction, and he often modified his plot and character development based on such feedback.
<p>
Dickens was regarded as the 'literary colossus' of his age. His 1843 novella, A Christmas Carol, is one of the most influential works ever written, and it remains popular and continues to inspire adaptations in every artistic genre. His creative genius has been praised by fellow writers—from Leo Tolstoy to G. K. Chesterton and George Orwell—for its realism, comedy, prose style, unique characterisations, and social criticism. On the other hand Oscar Wilde, Henry James and Virginia Woolf complained of a lack of psychological depth, loose writing, and a vein of saccharine sentimentalism.`,
	},
	{
		"Chuck Palahniuk", //TODO add more text to bio.
		"chuck_palahniuk.jpg",
		`Photo by <a href="http://www.flickr.com/photos/sourdiesel/484793142/">AlexRan</a>. CC-BY-SA`,
		"https://en.wikipedia.org/wiki/Chuck_Palahniuk",
		`Charles Michael "Chuck" Palahniuk (born February 21, 1962) is an American transgressional fiction novelist and freelance journalist. He is best known for the award-winning novel Fight Club, which was later made into a feature film. He maintains homes in the states of Oregon and Washington.`,
	},
	{
		"Cory Doctorow", //TODO add more text to bio.
		"cory_doctorow.jpg",
		`Photo by <a href="http://www.flickr.com/photos/null0/271975664">null0</a>. CC-BY-SA-2.0`,
		"https://en.wikipedia.org/wiki/Cory_Doctorow",
		`Cory Efram Doctorow (born July 17, 1971) is a Canadian-British blogger, journalist, and science fiction author who serves as co-editor of the weblog Boing Boing. He is an activist in favour of liberalising copyright laws and a proponent of the Creative Commons organization, using some of their licences for his books. Some common themes of his work include digital rights management, file sharing, and "post-scarcity" economics.`,
	},
	{
		"Dan Brown",
		"dan_brown.jpg",
		`Photo by Philip Scalia. CC-BY-SA-3.0`,
		"https://en.wikipedia.org/wiki/Dan_Brown",
		`Dan Brown (born June 22, 1964) is an American author of thriller fiction, best known for the 2003 bestselling novel, <cite>The Da Vinci Code</cite>. Brown's novels, which are treasure hunts set in a 24-hour time period, feature the recurring themes of cryptography, keys, symbols, codes, and conspiracy theories. His books have been translated into over 40 languages, and as of 2009, sold over 80 million copies. Two of them, <cite>The Da Vinci Code</cite> and <cite>Angels & Demons</cite>, have been adapted into feature films. The former opened amid great controversy and poor reviews, while the latter did only slightly better with critics.
<p>
Brown's novels that feature the lead character Robert Langdon also include historical themes and Christianity as recurring motifs, and as a result, have generated controversy. Brown states on his website that his books are not anti-Christian, though he is on a 'constant spiritual journey' himself, and says that his book <cite>The Da Vinci Code</cite> is simply "an entertaining story that promotes spiritual discussion and debate" and suggests that the book may be used "as a positive catalyst for introspection and exploration of our faith.`,
	},
	{
		"Daniel Defoe",
		"daniel_defoe.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Daniel_Defoe",
		`Daniel Defoe (ca. 1659–1661 to 24 April 1731), born Daniel Foe, was an English trader, writer, journalist, and pamphleteer, who gained fame for his novel <cite>Robinson Crusoe</cite>. Defoe is notable for being one of the earliest proponents of the novel, as he helped to popularise the form in Britain and along with others such as Richardson, is among the founders of the English novel. A prolific and versatile writer, he wrote more than 500 books, pamphlets and journals on various topics (including politics, crime, religion, marriage, psychology and the supernatural). He was also a pioneer of economic journalism.`,
	},
	{
		"David Foster Wallace",
		"david_foster_wallace.jpg",
		`Photo by <a href="http://www.flickr.com/photos/44124466908@N01/88166765">Steve Rhodes</a>. CC-BY 2.0.`,
		"https://en.wikipedia.org/wiki/David_Foster_Wallace",
		`David Foster Wallace (February 21, 1962 – September 12, 2008) was an award-winning American author of novels, essays, and short stories, and a professor at Pomona College in Claremont, California. He was widely known for his 1996 novel <cite>Infinite Jest</cite>. In 2005, Time magazine included the novel in its list of the 100 best English-language novels from 1923 to the present.
<p>
Los Angeles Times book editor David Ulin called Wallace "one of the most influential and innovative writers of the last 20 years". Wallace's unfinished novel, <cite>The Pale King</cite>, was published in 2011, and in 2012 was a finalist for the Pulitzer Prize.`,
	},
	{
		"Douglas Adams",
		"douglas_adams.jpg",
		`Photo by <a href="http://www.flickr.com/photos/michael_hughes/2944419184/">Michael Hughes</a>. CC-BY-SA-2.0`,
		"https://en.wikipedia.org/wiki/Douglas_Adams",
		`Douglas Noel Adams (11 March 1952 – 11 May 2001) was an English writer and dramatist. He is best known as the author of <cite>The Hitchhiker's Guide to the Galaxy</cite>, which started life in 1978 as a BBC radio comedy before developing into a "trilogy" of five books that sold over 15 million copies in his lifetime, a television series, several stage plays, comics, a computer game, and in 2005 a feature film. Adams's contribution to UK radio is commemorated in The Radio Academy's Hall of Fame.
<p>
Adams also wrote <cite>Dirk Gently's Holistic Detective Agency</cite> (1987) and <cite>The Long Dark Tea-Time of the Soul</cite> (1988), and co-wrote <cite>The Meaning of Liff</cite> (1983), <cite>Last Chance to See</cite> (1990), and three stories for the television series Doctor Who. A posthumous collection of his work, including an unfinished novel, was published as <cite>The Salmon of Doubt</cite> in 2002.
<p>
Adams became known as an advocate for environmental and conservation causes, and also as a lover of fast cars, cameras, and the Apple Macintosh. He was a staunch atheist, famously imagining a sentient puddle who wakes up one morning and thinks, "This is an interesting world I find myself in—an interesting hole I find myself in—fits me rather neatly, doesn't it? In fact it fits me staggeringly well, must have been made to have me in it!" Biologist Richard Dawkins dedicated his book, <cite>The God Delusion</cite> (2006), to Adams, writing on his death that, "Science has lost a friend, literature has lost a luminary, the mountain gorilla and the black rhino have lost a gallant defender."`,
	},
	{
		"Edgar Allan Poe",
		"edgar_allan_poe.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Edgar_Allan_Poe",
		`Edgar Allan Poe (born Edgar Poe; January 19, 1809 – October 7, 1849) was an American author, poet, editor and literary critic, considered part of the American Romantic Movement. Best known for his tales of mystery and the macabre, Poe was one of the earliest American practitioners of the short story and is considered the inventor of the detective fiction genre. He is further credited with contributing to the emerging genre of science fiction. He was the first well-known American writer to try to earn a living through writing alone, resulting in a financially difficult life and career.
<p>
Poe and his works influenced literature in the United States and around the world, as well as in specialized fields, such as cosmology and cryptography. Poe and his work appear throughout popular culture in literature, music, films, and television. A number of his homes are dedicated museums today.`,
	},
	{
		"Ernest Hemingway",
		"ernest_hemingway.jpg",
		"Photo by Lloyd Arnold. Public domain.",
		"https://en.wikipedia.org/wiki/Ernest_Hemingway",
		`Ernest Miller Hemingway (July 21, 1899 – July 2, 1961) was an American author and journalist. His economical and understated style had a strong influence on 20th-century fiction, while his life of adventure and his public image influenced later generations. Hemingway produced most of his work between the mid-1920s and the mid-1950s, and won the Nobel Prize in Literature in 1954. He published seven novels, six short story collections and two non-fiction works. Three novels, four collections of short stories and three non-fiction works were published posthumously. Many of these are considered classics of American literature.
<p>
Hemingway was raised in Oak Park, Illinois. After high school he reported for a few months for The Kansas City Star, before leaving for the Italian front to enlist with the World War I ambulance drivers. In 1918, he was seriously wounded and returned home. His wartime experiences formed the basis for his novel <cite>A Farewell to Arms</cite>. In 1922, he married Hadley Richardson, the first of his four wives. The couple moved to Paris, where he worked as a foreign correspondent, and fell under the influence of the modernist writers and artists of the 1920s "Lost Generation" expatriate community. <cite>The Sun Also Rises</cite>, Hemingway's first novel, was published in 1926.
<p>
After his 1927 divorce from Hadley Richardson, Hemingway married Pauline Pfeiffer. They divorced after he returned from the Spanish Civil War where he had acted as a journalist, and after which he wrote <cite>For Whom the Bell Tolls</cite>. Martha Gellhorn became his third wife in 1940. They separated when he met Mary Welsh in London during World War II; during which he was present at the Normandy Landings and liberation of Paris.
<p>
Shortly after the publication of <cite>The Old Man and the Sea</cite> in 1952, Hemingway went on safari to Africa, where he was almost killed in a plane crash that left him in pain or ill-health for much of the rest of his life. Hemingway had permanent residences in Key West, Florida, and Cuba during the 1930s and 1940s, but in 1959 he moved from Cuba to Ketchum, Idaho, where he committed suicide in the summer of 1961.`,
	},
	{
		"George Orwell",
		"george_orwell.jpg",
		"", // Public domain, anonymous photographer.
		"https://en.wikipedia.org/wiki/George_Orwell",
		`Eric Arthur Blair (25 June 1903 – 21 January 1950), better known by his pen name George Orwell, was an English novelist and journalist. His work is marked by keen intelligence and wit, a profound awareness of social injustice, an intense opposition to totalitarianism, a passion for clarity in language and a belief in democratic socialism.
<p>
Considered perhaps the 20th century's best chronicler of English culture, Orwell wrote literary criticism, poetry, fiction and polemical journalism. He is best known for the dystopian novel <cite>Nineteen Eighty-Four</cite> (1949) and the allegorical novella <cite>Animal Farm</cite> (1945), which together have sold more copies than any two books by any other 20th-century author. His book <cite>Homage to Catalonia</cite> (1938), an account of his experiences in the Spanish Civil War, is widely acclaimed, as are his numerous essays on politics, literature, language and culture. In 2008, The Times ranked him second on a list of "The 50 greatest British writers since 1945".
<p>
Orwell's influence on popular and political culture endures, and several of his neologisms, along with the term Orwellian — a byword for totalitarian or manipulative social practices — have entered the vernacular.`,
	},
	{
		"Gertrude Stein", //TODO add more text to bio.
		"gertrude_stein.jpg",
		"", // Public domain.
		"https://en.wikipedia.org/wiki/Gertrude_Stein",
		`Gertrude Stein (February 3, 1874 – July 27, 1946) was an American writer, poet, and art collector who spent most of her life in France.`,
	},
	{
		"H. G. Wells", //TODO mention his famous books in bio.
		"h_g_wells.jpg",
		"", // Public domain.
		"https://en.wikipedia.org/wiki/H._G._Wells",
		`Herbert George "H.G." Wells (21 September 1866 – 13 August 1946) was a British author, now best known for his work in the science fiction genre. He was also a prolific writer in many other genres, including contemporary novels, history, politics and social commentary, even writing text books and rules for war games. Together with Jules Verne and Hugo Gernsback, Wells has been referred to as "The Father of Science Fiction".
<p>
Wells's earliest specialized training was in biology, and his thinking on ethical matters took place in a specifically and fundamentally Darwinian context. He was also from an early date an outspoken socialist, often (but not always, as the beginning of the First World War) sympathizing with pacifist views. His later works became increasingly political and didactic, and he sometimes indicated on official documents that his profession was that of "Journalist". Many of his novels, particularly those of his middle period (1900–1920), had nothing to do with science fiction. They described lower-middle class life (<cite>Kipps</cite>; <cite>The History of Mr Polly</cite>), and this work sometimes led Wells to be touted as a worthy successor to Charles Dickens.`,
	},
	{
		"H. P. Lovecraft",
		"h_p_lovecraft.jpg",
		"", // Public domain, 1915
		"https://en.wikipedia.org/wiki/H._P._Lovecraft",
		`Howard Phillips Lovecraft (August 20, 1890 – March 15, 1937) — known as
H. P. Lovecraft — was an American author of horror, fantasy and science fiction,
especially the subgenre known as weird fiction.
<p>
Lovecraft's guiding aesthetic and philosophical principle was what he termed
"cosmicism" or "cosmic horror", the idea that life is incomprehensible to human
minds and that the universe is fundamentally inimical to the interests of
humankind. As such, his stories express a profound indifference to human beliefs
and affairs. Lovecraft is best known for his Cthulhu Mythos story cycle and the
<cite>Necronomicon</cite>, a fictional grimoire of magical rites and forbidden lore.
<p>
Although Lovecraft's readership was limited during his lifetime, his reputation
has grown over the decades, and he is now regarded as one of the most
influential horror writers of the 20th century. According to Joyce Carol Oates,
Lovecraft—as with Edgar Allan Poe in the 19th century—has exerted "an
incalculable influence on succeeding generations of writers of horror fiction".
Stephen King called Lovecraft "the twentieth century's greatest practitioner of
the classic horror tale." King has even made it clear in his
semi-autobiographical non-fiction book <cite>Danse Macabre</cite> that Lovecraft was
responsible for King's own fascination with horror and the macabre, and was the
single largest figure to influence his fiction writing. His stories have also
been adapted into plays, films and games.`,
	},
	{
		"Harry Harrison", //TODO add more text to bio.
		"harry_harrison.jpg",
		`Photo by <a href="http://home.agh.edu.pl/~szymon/">Szymon Sokół</a>. CC-BY-SA 3.0`,
		"https://en.wikipedia.org/wiki/Harry_Harrison",
		`Harry Harrison (born March 12, 1925) is an American science fiction author best known for his character the Stainless Steel Rat and the novel <cite>Make Room! Make Room!</cite> (1966), the basis for the film <cite>Soylent Green</cite> (1973). He is also (with Brian Aldiss) co-president of the Birmingham Science Fiction Group.`,
	},
	{
		"Ian Fleming",
		"", // CC or public domain photos not found :(
		"",
		"https://en.wikipedia.org/wiki/Ian_Fleming",
		`Ian Lancaster Fleming (28 May 1908 – 12 August 1964) was an English author, journalist and Naval Intelligence Officer. Fleming is best known for creating the fictional spy James Bond and the series of twelve novels and nine short stories about the character. Fleming was from a wealthy family, connected to the merchant bank Robert Fleming & Co. and his father was MP for Henley from 1910 until his death on the Western Front in 1917. Educated at Eton, the Royal Military Academy, Sandhurst and the universities of Munich and Geneva, Fleming moved through a number of jobs before he started writing.
<p>
The Bond books are among the biggest-selling series of fictional books of all time, having sold over 100 million copies worldwide. Fleming also wrote the children's story <cite>Chitty-Chitty-Bang-Bang</cite> and two works of non-fiction. While working in British Naval Intelligence during the Second World War, Fleming was involved in the planning stages of Operation Mincemeat and Operation Golden Eye, the former of which was successfully carried out. Fleming was also involved in the planning and overseeing of two active service units, 30 Assault Unit and T-Force.
<p>
His experiences of the people he met during his wartime service provided much of the background and detail of the Bond novels and his career as a journalist added colour and depth to the stories. In 2008, The Times ranked Fleming fourteenth on its list of "The 50 greatest British writers since 1945".`,
	},
	{
		"Isaac Asimov",
		"isaac_asimov.jpg",
		"", // Public domain, see http://en.wikipedia.org/wiki/File:Isaac.Asimov01.jpg
		"https://en.wikipedia.org/wiki/Isaac_Asimov",
		`Isaac Asimov (born Isaak Yudovich Ozimov, c. January 2, 1920 – April 6, 1992) was an American author and professor of biochemistry at Boston University, best known for his works of science fiction and for his popular science books. Asimov was one of the most prolific writers of all time, having written or edited more than 500 books and an estimated 90,000 letters and postcards. His works have been published in all ten major categories of the Dewey Decimal System (although his only work in the 100s—which covers philosophy and psychology—was a foreword for The Humanist Way).
<p>
Asimov is widely considered a master of hard science fiction and, along with Robert A. Heinlein and Arthur C. Clarke, he was considered one of the "Big Three" science fiction writers during his lifetime. Asimov's most famous work is the <cite>Foundation Series</cite>; his other major series are the <cite>Galactic Empire</cite> series and the <cite>Robot</cite> series, both of which he later tied into the same fictional universe as the <cite>Foundation Series</cite> to create a unified "future history" for his stories much like those pioneered by Robert A. Heinlein and previously produced by Cordwainer Smith and Poul Anderson. He wrote many short stories, among them <cite>Nightfall</cite>, which in 1964 was voted by the Science Fiction Writers of America the best short science fiction story of all time. Asimov wrote the <cite>Lucky Starr</cite> series of juvenile science-fiction novels using the pen name Paul French.
<p>
The prolific Asimov also wrote mysteries and fantasy, as well as much non-fiction. Most of his popular science books explain scientific concepts in a historical way, going as far back as possible to a time when the science in question was at its simplest stage. He often provides nationalities, birth dates, and death dates for the scientists he mentions, as well as etymologies and pronunciation guides for technical terms. Examples include <cite>Guide to Science</cite>, the three volume set <cite>Understanding Physics</cite>, <cite>Asimov's Chronology of Science and Discovery</cite>, as well as works on astronomy, mathematics, the Bible, William Shakespeare's writing and chemistry.
<p>
The asteroid 5020 Asimov, a crater on the planet Mars, a Brooklyn, New York elementary school, and one Isaac Asimov literary award are named in his honor.`,
	},
	{
		"J. D. Salinger",
		"", // CC or public domain photos not found :(
		"",
		"http://en.wikipedia.org/wiki/J._D._Salinger",
		`Jerome David "J. D." Salinger (January 1, 1919 – January 27, 2010) was an American author, best known for his only novel, The Catcher in the Rye (1951), and his reclusive nature. His last original published work was in 1965; he gave his last interview in 1980.`,
	},
	{
		"J. K. Rowling",
		"j_k_rowling.jpg",
		`Photo by <a href="http://www.flickr.com/photos/fast50">Daniel Ogren</a>. CC-BY 2.0`,
		"https://en.wikipedia.org/wiki/J._K._Rowling",
		`Joanne "Jo" Rowling, (born 31 July 1965), known as J. K. Rowling, is a British novelist, best known as the author of the <cite>Harry Potter</cite> fantasy series. The Potter books have gained worldwide attention, won multiple awards, sold more than 400 million copies to become the best-selling book series in history and been the basis for a popular series of films, in which Rowling had overall approval on the scripts as well as maintaining creative control by serving as a producer on the final installment. Rowling conceived the idea for the series on a train trip from Manchester to London in 1990.
<p>
Rowling progressed from living on social security to multi-millionaire status within five years. As of March 2011, when its latest world billionaires list was published, Forbes estimated Rowling's net worth to be US$1 billion. The 2008 Sunday Times Rich List estimated Rowling's fortune at £560 million ($798 million), ranking her as the twelfth richest woman in the United Kingdom. Forbes ranked Rowling as the forty-eighth most powerful celebrity of 2007, and Time magazine named her as a runner-up for its 2007 Person of the Year, noting the social, moral, and political inspiration she has given her fans. In October 2010, J. K. Rowling was named 'Most Influential Woman in Britain' by leading magazine editors. She has become a notable philanthropist.`,
	},
	{
		"J. R. R. Tolkien",
		"j_r_r_tolkien.jpg",
		"", // Public domain, 1916.
		"https://en.wikipedia.org/wiki/J._R._R._Tolkien",
		`John Ronald Reuel Tolkien, CBE (3 January 1892 – 2 September 1973) was an English writer, poet, philologist, and university professor, best known as the author of the classic high fantasy works <cite>The Hobbit</cite>, <cite>The Lord of the Rings</cite>, and <cite>The Silmarillion</cite>.
<p>
Tolkien was Rawlinson and Bosworth Professor of Anglo-Saxon at Pembroke College, Oxford, from 1925 to 1945 and Merton Professor of English Language and Literature there from 1945 to 1959. He was a close friend of C. S. Lewis — they were both members of the informal literary discussion group known as the Inklings. Tolkien was appointed a Commander of the Order of the British Empire by Queen Elizabeth II on 28 March 1972.
<p>
After his death, Tolkien's son Christopher published a series of works based on his father's extensive notes and unpublished manuscripts, including <cite>The Silmarillion</cite>. These, together with <cite>The Hobbit</cite> and <cite>The Lord of the Rings</cite> form a connected body of tales, poems, fictional histories, invented languages, and literary essays about a fantasy world called Arda, and Middle-earth within it. Between 1951 and 1955, Tolkien applied the term legendarium to the larger part of these writings.
<p>
While many other authors had published works of fantasy before Tolkien, the great success of <cite>The Hobbit</cite> and <cite>The Lord of the Rings</cite> led directly to a popular resurgence of the genre. This has caused Tolkien to be popularly identified as the "father" of modern fantasy literature. In 2008, The Times ranked him sixth on a list of "The 50 greatest British writers since 1945". Forbes ranked him the 5th top-earning dead celebrity in 2009.`,
	},
	{
		"Jack London",
		"jack_london.jpg",
		"", // Public domain, 1905
		"https://en.wikipedia.org/wiki/Jack_London",
		`John Griffith "Jack" London (born John Griffith Chaney, January 12, 1876 – November 22, 1916) was an American author, journalist, and social activist. He was a pioneer in the then-burgeoning world of commercial magazine fiction and was one of the first fiction writers to obtain worldwide celebrity and a large fortune from his fiction alone. He is best remembered as the author of <cite>Call of the Wild</cite> and <cite>White Fang</cite>, both set in the Klondike Gold Rush. He also wrote of the South Pacific in such stories as "The Pearls of Parlay" and "The Heathen", and of the San Francisco Bay area in The Sea Wolf.
<p>
London was a passionate advocate of unionization, socialism, and the rights of workers and wrote several powerful works dealing with these topics such as his dystopian novel, <cite>The Iron Heel</cite> and his non-fiction exposé, <cite>The People of the Abyss</cite>.`,
	},
	{
		"James Fenimore Cooper",
		"james_fenimore_cooper.jpg",
		"", // Public domain.
		"https://en.wikipedia.org/wiki/James_Fenimore_Cooper",
		`James Fenimore Cooper (September 15, 1789 – September 14, 1851) was a prolific and popular American writer of the early 19th century. He is best remembered as a novelist who wrote numerous sea-stories and the historical novels known as the <cite>Leatherstocking Tales</cite>, featuring frontiersman Natty Bumppo. Among his most famous works is the Romantic novel <cite>The Last of the Mohicans</cite>, often regarded as his masterpiece.`,
	},
	{
		"James Joyce",
		"james_joyce.jpg",
		"", // Public domain, 1915.
		"https://en.wikipedia.org/wiki/James_Joyce",
		`James Augustine Aloysius Joyce (2 February 1882 – 13 January 1941) was an Irish novelist and poet, considered to be one of the most influential writers in the modernist avant-garde of the early 20th century.
<p>
Joyce is best known for <cite>Ulysses</cite> (1922), a landmark work in which the episodes of Homer's Odyssey are paralleled in an array of contrasting literary styles, perhaps most prominently the stream of consciousness technique he perfected. Other major works are the short-story collection <cite>Dubliners</cite> (1914), and the novels <cite>A Portrait of the Artist as a Young Man</cite> (1916) and <cite>Finnegans Wake</cite> (1939). His complete oeuvre includes three books of poetry, a play, occasional journalism, and his published letters.`,
	},
	{
		"Jane Austen",
		"jane_austen.jpg",
		"", // Public domain.
		"https://en.wikipedia.org/wiki/Jane_Austen",
		`Jane Austen (16 December 1775 – 18 July 1817) was an English novelist whose works of romantic fiction, set among the landed gentry, earned her a place as one of the most widely read writers in English literature. Her realism and biting social commentary has gained her historical importance among scholars and critics.
<p>
Austen lived her entire life as part of a close-knit family located on the lower fringes of the English landed gentry. She was educated primarily by her father and older brothers as well as through her own reading. The steadfast support of her family was critical to her development as a professional writer. Her artistic apprenticeship lasted from her teenage years into her thirties. During this period, she experimented with various literary forms, including the epistolary novel which she tried then abandoned, and wrote and extensively revised three major novels and began a fourth. From 1811 until 1816, with the release of <cite>Sense and Sensibility</cite> (1811), <cite>Pride and Prejudice</cite> (1813), <cite>Mansfield Park</cite> (1814) and <cite>Emma</cite> (1816), she achieved success as a published writer. She wrote two additional novels, <cite>Northanger Abbey</cite> and <cite>Persuasion</cite>, both published posthumously in 1818, and began a third, which was eventually titled <cite>Sanditon</cite>, but died before completing it.
<p>
Austen's works critique the novels of sensibility of the second half of the 18th century and are part of the transition to 19th-century realism. Her plots, though fundamentally comic, highlight the dependence of women on marriage to secure social standing and economic security. Her work brought her little personal fame and only a few positive reviews during her lifetime, but the publication in 1869 of her nephew's <cite>A Memoir of Jane Austen</cite> introduced her to a wider public, and by the 1940s she had become widely accepted in academia as a great English writer. The second half of the 20th century saw a proliferation of Austen scholarship and the emergence of a Janeite fan culture.`,
	},
	{
		"Jonathan Swift",
		"jonathan_swift.jpg",
		"", // Public domain.
		"https://en.wikipedia.org/wiki/Jonathan_Swift",
		`Jonathan Swift (30 November 1667 – 19 October 1745) was an Anglo-Irish satirist, essayist, political pamphleteer (first for the Whigs, then for the Tories), poet and cleric who became Dean of St. Patrick's Cathedral, Dublin.
<p>
He is remembered for works such as <cite>Gulliver's Travels</cite>, <cite>A Modest Proposal</cite>, <cite>A Journal to Stella</cite>, <cite>Drapier's Letters</cite>, <cite>The Battle of the Books</cite>, <cite>An Argument Against Abolishing Christianity</cite>, and <cite>A Tale of a Tub</cite>. Swift is probably the foremost prose satirist in the English language, and is less well known for his poetry. Swift originally published all of his works under pseudonyms, such as Lemuel Gulliver, Isaac Bickerstaff, M.B. Drapier, or anonymously. He is also known for being a master of two styles of satire: the Horatian and Juvenalian styles.`,
	},
	{
		"Kurt Vonnegut",
		"", // No image.
		"",
		"https://en.wikipedia.org/wiki/Kurt_Vonnegut",
		`Kurt Vonnegut, Jr. (November 11, 1922 – April 11, 2007) was a 20th century American writer. His works such as <cite>Cat's Cradle</cite> (1963), <cite>Slaughterhouse-Five</cite> (1969), and <cite>Breakfast of Champions</cite> (1973) blend satire, gallows humor, and science fiction. As a citizen he was a lifelong supporter of the American Civil Liberties Union and a critical leftist intellectual. He was known for his humanist beliefs and was honorary president of the American Humanist Association.`,
	},
	{
		"L. Frank Baum",
		"l_frank_baum.jpg",
		"", // Public domain, 1911.
		"https://en.wikipedia.org/wiki/L._Frank_Baum",
		`Lyman "L." Frank Baum (May 15, 1856 – May 6, 1919) was an American author of children's books, best known for writing <cite>The Wonderful Wizard of Oz</cite>. He wrote thirteen novel sequels, nine other fantasy novels, and a host of other works (55 novels in total, plus four "lost" novels, 82 short stories, over 200 poems, an unknown number of scripts, and many miscellaneous writings), and made numerous attempts to bring his works to the stage and screen.`,
	},
	{
		"Leo Tolstoy",
		"leo_tolstoy.jpg",
		"", // Public domain, 1908
		"https://en.wikipedia.org/wiki/Leo_Tolstoy",
		`Leo Tolstoy, was a Russian writer who is regarded as one of the greatest authors of all time.

Born to an aristocratic Russian family in 1828, he is best known for the novels <cite>War and Peace</cite> (1869) and <cite>Anna Karenina</cite> (1877), often cited as pinnacles of realist fiction. He first achieved literary acclaim in his twenties with his semi-autobiographical trilogy, <cite>Childhood, Boyhood, and Youth</cite> (1852–1856), and <cite>Sevastopol Sketches</cite> (1855), based upon his experiences in the Crimean War. Tolstoy's fiction includes dozens of short stories and several novellas such as <cite>The Death of Ivan Ilyich</cite>, <cite>Family Happiness</cite>, and <cite>Hadji Murad</cite>. He also wrote plays and numerous philosophical essays.

In the 1870s Tolstoy experienced a profound moral crisis, followed by what he regarded as an equally profound spiritual awakening, as outlined in his non-fiction work <cite>A Confession</cite>. His literal interpretation of the ethical teachings of Jesus, centering on the Sermon on the Mount, caused him to become a fervent Christian anarchist and pacifist. Tolstoy's ideas on nonviolent resistance, expressed in such works as <cite>The Kingdom of God Is Within You</cite>, were to have a profound impact on such pivotal 20th-century figures as Mohandas Gandhi, Martin Luther King, Jr., and James Bevel. Tolstoy also became a dedicated advocate of Georgism, the economic philosophy of Henry George, which he incorporated into his writing, particularly <cite>Resurrection</cite>.`,
	},
	{
		"Lewis Carroll",
		"lewis_carroll.jpg",
		"", // Public domain, 1898
		"https://en.wikipedia.org/wiki/Lewis_Carroll",
		`Charles Lutwidge Dodgson (27 January 1832 – 14 January 1898), better known by the pseudonym Lewis Carroll, was an English author, mathematician, logician, Anglican deacon and photographer. His most famous writings are <cite>Alice's Adventures in Wonderland</cite> and its sequel <cite>Through the Looking-Glass</cite>, as well as the poems "The Hunting of the Snark" and "Jabberwocky", all examples of the genre of literary nonsense. He is noted for his facility at word play, logic, and fantasy.`,
	},
	{
		"Margaret Atwood",
		"margaret_atwood.jpg",
		"Photo by Vanwaffle. CC-BY-SA 3.0",
		"https://en.wikipedia.org/wiki/Margaret_Atwood",
		`Margaret Eleanor Atwood, CC OOnt FRSC (born November 18, 1939) is a Canadian poet, novelist, literary critic, essayist, and environmental activist. She is among the most-honored authors of fiction in recent history. She is a winner of the Arthur C. Clarke Award and Prince of Asturias award for Literature, has been shortlisted for the Booker Prize five times, winning once, and has been a finalist for the Governor General's Award seven times, winning twice.
<p>
While she is best known for her work as a novelist, she is also a poet, having published 15 books of poetry to date. Many of her poems have been inspired by myths and fairy tales, which have been interests of hers from an early age. Atwood has published short stories in Tamarack Review, Alphabet, Harper's, CBC Anthology, Ms., Saturday Night, and many other magazines. She has also published four collections of stories and three collections of unclassifiable short prose works.`,
	},
	{
		"Margaret Mitchell",
		"margaret_mitchell.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Margaret_Mitchell",
		`Margaret Munnerlyn Mitchell (November 8, 1900 – August 16, 1949) was an American author and journalist. One novel by Mitchell was published during her lifetime, the American Civil War-era novel, <cite>Gone with the Wind</cite>. For it she won the National Book Award for Most Distinguished Novel of 1936 and the Pulitzer Prize for Fiction in 1937.
<p>
In more recent years, a collection of Mitchell's girlhood writings and a novella she wrote as a teenager, <cite>Lost Laysen</cite>, have been published. A collection of articles written by Mitchell for <cite>The Atlanta Journal</cite> was republished in book form.`,
	},
	{
		"Mario Puzo",
		"", // No free image.
		"",
		"https://en.wikipedia.org/wiki/Mario_Puzo",
		`Mario Gianluigi Puzo (October 15, 1920 – July 2, 1999) was an Italian American author and screenwriter, known for his novels about the Mafia, including <cite>The Godfather</cite> (1969), which he later co-adapted into a film by Francis Ford Coppola. He won the Academy Award for Best Adapted Screenplay in both 1972 and 1974.`,
	},
	{
		"Mark Twain",
		"mark_twain.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Mark_Twain",
		`Samuel Langhorne Clemens (November 30, 1835 – April 21, 1910), better known by his pen name Mark Twain, was an American author and humorist. He is most noted for his novels, <cite>The Adventures of Tom Sawyer</cite> (1876), and its sequel, <cite>Adventures of Huckleberry Finn</cite> (1885), the latter often called "the Great American Novel."
<p>
Twain grew up in Hannibal, Missouri, which would later provide the setting for <cite>Huckleberry Finn</cite> and <cite>Tom Sawyer</cite>. He apprenticed with a printer. He also worked as a typesetter and contributed articles to his older brother Orion's newspaper. After toiling as a printer in various cities, he became a master riverboat pilot on the Mississippi River, before heading west to join Orion. He was a failure at gold mining, so he next turned to journalism. While a reporter, he wrote a humorous story, "The Celebrated Jumping Frog of Calaveras County", which became very popular and brought nationwide attention. His travelogues were also well-received. Twain had found his calling.
<p>
He achieved great success as a writer and public speaker. His wit and satire earned praise from critics and peers, and he was a friend to presidents, artists, industrialists, and European royalty.
<p>
He lacked financial acumen, and, though he made a great deal of money from his writings and lectures, he squandered it on various ventures, in particular the Paige Compositor, and was forced to declare bankruptcy. With the help of Henry Huttleston Rogers he eventually overcame his financial troubles. Twain worked hard to ensure that all of his creditors were paid in full, even though his bankruptcy had relieved him of the legal responsibility.
<p>
He was lauded as the "greatest American humorist of his age," and William Faulkner called Twain "the father of American literature."`,
	},
	{
		"Mary Shelley",
		"mary_shelley.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Mary_Shelley",
		`Mary Shelley (née Mary Wollstonecraft Godwin; 30 August 1797 – 1 February 1851) was an English novelist, short story writer, dramatist, essayist, biographer, and travel writer, best known for her Gothic novel <cite>Frankenstein</cite>: or, <cite>The Modern Prometheus</cite> (1818). She also edited and promoted the works of her husband, the Romantic poet and philosopher Percy Bysshe Shelley. Her father was the political philosopher William Godwin, and her mother was the Philosopher and feminist Mary Wollstonecraft.`,
	},
	{
		"Neil Gaiman",
		"neil_gaiman.jpg",
		`Photo by <a href="https://commons.wikimedia.org/wiki/User:KyleCassidy">Kyle Cassidy</a>. CC-BY-SA 3.0`,
		"https://en.wikipedia.org/wiki/Neil_Gaiman",
		`Neil Richard Gaiman (born 10 November 1960) is an English author of short fiction, novels, comic books, graphic novels, audio theatre and films. His notable works include the comic book series <cite>The Sandman</cite> and novels <cite>Stardust</cite>, <cite>American Gods</cite>, <cite>Coraline</cite>, and <cite>The Graveyard Book</cite>. He has won numerous awards, including Hugo, Nebula, Bram Stoker, Newbery Medal, and Carnegie Medal in Literature. He is the first author to win both the Newbery and the Carnegie medals for the same work.`,
	},
	{
		"Oscar Wilde",
		"oscar_wilde.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Oscar_Wilde",
		`Oscar Fingal O'Flahertie Wills Wilde (16 October 1854 – 30 November 1900) was an Irish writer and poet. After writing in different forms throughout the 1880s, he became one of London's most popular playwrights in the early 1890s. Today he is remembered for his epigrams and plays, and the circumstances of his imprisonment which was followed by his early death.`,
	},
	{
		"P. G. Wodehouse",
		"p_g_wodehouse.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/P._G._Wodehouse",
		`Sir Pelham Grenville Wodehouse, KBE (15 October 1881 – 14 February 1975) was an English humorist, whose body of work includes novels, short stories, plays, poems, song lyrics and numerous pieces of journalism. He enjoyed enormous popular success during a career that lasted more than seventy years and his many writings continue to be widely read. Despite the political and social upheavals that occurred during his life, much of which was spent in France and the United States, Wodehouse's main canvas remained that of a pre- and post-World War I English upper class society, reflecting his birth, education and youthful writing career.
<p>
An acknowledged master of English prose, Wodehouse has been admired both by contemporaries such as Hilaire Belloc, Evelyn Waugh and Rudyard Kipling and by recent writers such as Stephen Fry, Douglas Adams, J. K. Rowling, and John Le Carré.
<p>
Best known today for the Jeeves and Blandings Castle novels and short stories, Wodehouse was also a playwright and lyricist who was part author and writer of 15 plays and of 250 lyrics for some 30 musical comedies, many of them produced in collaboration with Jerome Kern and Guy Bolton. He worked with Cole Porter on the musical <cite>Anything Goes</cite> (1934), wrote the lyrics for the hit song "Bill" in Kern's <cite>Show Boat</cite> (1927), wrote lyrics to Sigmund Romberg's music for the Gershwin – Romberg musical <cite>Rosalie</cite> (1928) and collaborated with Rudolf Friml on a musical version of <cite>The Three Musketeers</cite> (1928). He is in the Songwriters Hall of Fame.`,
	},
	{
		"Ray Bradbury",
		"ray_bradbury.jpg",
		`Photo by <a href="http://www.flickr.com/people/alan-light/">Alan Light</a>. CC-BY 2.0`,
		"https://en.wikipedia.org/wiki/Ray_Bradbury",
		`Ray Douglas Bradbury (August 22, 1920 – June 5, 2012) was an American fantasy, science fiction, horror and mystery fiction writer. Best known for his dystopian novel <cite>Fahrenheit 451</cite> (1953) and for the science fiction and horror stories gathered together as <cite>The Martian Chronicles</cite> (1950) and <cite>The Illustrated Man</cite> (1951), Bradbury was one of the most celebrated 20th-century American writers. Many of Bradbury's works have been adapted into comic books, television shows and films.`,
	},
	{
		"Raymond Chandler",
		"", // No free photo
		"",
		"https://en.wikipedia.org/wiki/Raymond_Chandler",
		`Raymond Thornton Chandler (July 23, 1888 – March 26, 1959) was an American novelist and screenwriter.
<p>
In 1932, at age forty-four, Raymond Chandler decided to become a detective fiction writer after losing his job as an oil company executive during the Depression. His first short story, "Blackmailers Don't Shoot", was published in 1933 in <cite>Black Mask</cite>, a popular pulp magazine. His first novel, <cite>The Big Sleep</cite>, was published in 1939. In addition to his short stories, Chandler published just seven full novels during his lifetime (though an eighth in progress at his death was completed by Robert B. Parker). All but <cite>Playback</cite> have been realized into motion pictures, some several times. In the year before he died, he was elected president of the Mystery Writers of America.
<p>
Chandler had an immense stylistic influence on American popular literature, and is considered by many to be a founder, along with Dashiell Hammett, James M. Cain and other <cite>Black Mask</cite> writers, of the hard-boiled school of detective fiction. His protagonist, Philip Marlowe, along with Hammett's Sam Spade, is considered by some to be synonymous with "private detective," both having been played on screen by Humphrey Bogart, whom many considered to be the quintessential Marlowe.`,
	},
	{
		"Robert Louis Stevenson",
		"robert_louis_stevenson.jpg",
		"", // Public domain
		"https://en.wikipedia.org/wiki/Robert_Louis_Stevenson",
		`Robert Louis Balfour Stevenson (13 November 1850 – 3 December 1894) was a Scottish novelist, poet, essayist, and travel writer. His most famous works are <cite>Treasure Island</cite>, <cite>Kidnapped</cite>, and <cite>Strange Case of Dr Jekyll and Mr Hyde</cite>.
<p>
A literary celebrity during his lifetime, Stevenson now ranks among the 26 most translated authors in the world. His works have been admired by many other writers, including Jorge Luis Borges, Ernest Hemingway, Rudyard Kipling, Marcel Schwob, Vladimir Nabokov, J. M. Barrie, and G. K. Chesterton, who said of him that he "seemed to pick the right word up on the point of his pen, like a man playing spillikins."`,
	},
	{
		"Rudyard Kipling",
		"rudyard_kipling.jpg",
		"", // Public domain, 1912
		"https://en.wikipedia.org/wiki/Rudyard_Kipling",
		`Joseph Rudyard Kipling (30 December 1865 – 18 January 1936) was an English short-story writer, poet, and novelist chiefly remembered for his tales and poems of British soldiers in India, and his tales for children.
<p>
He was born in Bombay, in the Bombay Presidency of British India, and was taken by his family to England when he was five years old. Kipling is best known for his works of fiction, including <cite>The Jungle Book</cite> (a collection of stories which includes "Rikki-Tikki-Tavi"), <cite>Just So Stories</cite> (1902), <cite>Kim</cite> (1901) (a tale of adventure), many short stories, including "The Man Who Would Be King" (1888); and his poems, including "Mandalay" (1890), "Gunga Din" (1890), "The White Man's Burden" (1899) and "If—" (1910). He is regarded as a major "innovator in the art of the short story"; his children's books are enduring classics of children's literature; and his best works are said to exhibit "a versatile and luminous narrative gift".
<p>
Kipling was one of the most popular writers in England, in both prose and verse, in the late 19th and early 20th centuries. Henry James said: "Kipling strikes me personally as the most complete man of genius (as distinct from fine intelligence) that I have ever known." In 1907 he was awarded the Nobel Prize in Literature, making him the first English-language writer to receive the prize, and to date he remains its youngest recipient. Among other honours, he was sounded out for the British Poet Laureateship and on several occasions for a knighthood, all of which he declined.`,
	},
	{
		"Stephen King",
		"stephen_king.jpg",
		`Photo by <a href="http://www.flickr.com/photos/pinguino/409180680/">Pinguino</a>. CC-BY 2.0`,
		"https://en.wikipedia.org/wiki/Stephen_King",
		`Stephen Edwin King (born September 21, 1947) is an American author of contemporary horror, suspense, science fiction and fantasy. His books have sold more than 350 million copies and have been adapted into a number of feature films, television movies and comic books. King has published 50 novels, including seven under the pen-name of Richard Bachman, and five non-fiction books. He has written nearly two hundred short stories, most of which have been collected in nine collections of short fiction. Many of his stories are set in his home state of Maine.
<p>
King has received Bram Stoker Awards, World Fantasy Awards, British Fantasy Society Awards, his novella <cite>The Way Station</cite> was a Nebula Award novelette nominee, and his short story "The Man in the Black Suit" received the O. Henry Award. In 2003, the National Book Foundation awarded him the Medal for Distinguished Contribution to American Letters. He has also received awards for his contribution to literature for his whole career, such as the World Fantasy Award for Life Achievement (2004), the Canadian Booksellers Association Lifetime Achievement Award (2007) and the Grand Master Award from the Mystery Writers of America (2007).`,
	},
	{
		"Stephenie Meyer",
		"stephenie_meyer.jpg",
		`Photo by <a href="http://en.wikipedia.org/wiki/File:Stephenie_Meyer_by_Gage_Skidmore.jpg">Gage Skidmore</a>. CC-BY-SA 3.0`,
		"https://en.wikipedia.org/wiki/Stephenie_Meyer",
		`Stephenie Meyer (née Morgan; born December 24, 1973) is an American young adult author and producer, best known for her vampire romance series <cite>Twilight</cite>. The Twilight novels have gained worldwide recognition and sold over 100 million copies, with translations into 37 different languages. Meyer was the bestselling author of 2008 and 2009 in America, having sold over 29 million books in 2008, and 26.5 million books in 2009. Twilight was the best-selling book of 2008 in US bookstores.
<p>
Meyer was ranked #49 on Time magazine's list of the "100 Most Influential People in 2008", and was included in the Forbes Celebrity 100 list of the world's most powerful celebrities in 2009, entering at #26. Her annual earnings exceeded $50 million. In 2010, Forbes ranked her as the #59 most powerful celebrity with annual earnings of $40 million.`,
	},
	{
		"Ursula K. Le Guin",
		"ursula_k_le_guin.jpg",
		`Photo by Hajor. CC-BY-SA 1.0`,
		"https://en.wikipedia.org/wiki/Ursula_K._Le_Guin",
		`Ursula Kroeber Le Guin (born October 21, 1929) is an American author of novels, children's books, and short stories, mainly in the genres of fantasy and science fiction. She has also written poetry and essays.
First published in the 1960s, her work has often depicted futuristic or imaginary worlds alternative to our own in politics, natural environment, gender, religion, sexuality and ethnography.
<p>
She has won the Hugo Award, Nebula Award, Locus Award, and World Fantasy Award several times each.`,
	},
	{
		"Vladimir Nabokov",
		"", // No free image
		"",
		"https://en.wikipedia.org/wiki/Vladimir_Nabokov",
		`Vladimir Vladimirovich Nabokov (Russian: Влади́мир Влади́мирович Набо́ков; 22 April 1899 – 2 July 1977) was a Russian-American novelist. Nabokov's first nine novels were in Russian. He then rose to international prominence as a writer of English prose. He also made serious contributions as a lepidopterist and chess composer.
<p>
Nabokov's <cite>Lolita</cite> (1955) is his most famous novel, and often considered his finest work in English. It exhibits the love of intricate word play and synesthetic detail that characterised all his works. The novel was ranked at No. 4 in the list of the Modern Library 100 Best Novels. <cite>Pale Fire</cite> (1962) was ranked at No. 53 on the same list. His memoir, Speak, Memory, was listed No. 8 on the Modern Library nonfiction list. He was a finalist for the National Book Award for Fiction seven times, but never won it.`,
	},
	{
		"William Gibson",
		"william_gibson.jpg",
		`Photo by <a href="http://flickr.com/photos/45853044@N00">Gonzo Bonzo</a>. CC-BY-SA 2.0`,
		"https://en.wikipedia.org/wiki/William_Gibson",
		`William Ford Gibson (born March 17, 1948) is an American-Canadian speculative fiction novelist who has been called the "noir prophet" of the cyberpunk subgenre. Gibson coined the term "cyberspace" in his short story "Burning Chrome" (1982) and later popularized the concept in his debut novel, <cite>Neuromancer</cite> (1984). In envisaging cyberspace, Gibson created an iconography for the information age before the ubiquity of the Internet in the 1990s. He is also credited with predicting the rise of reality television and with establishing the conceptual foundations for the rapid growth of virtual environments such as video games and the World Wide Web.
<p>
 After expanding on <cite>Neuromancer</cite> with two more novels to complete the dystopic Sprawl trilogy, Gibson became an important author of another science fiction sub-genre—steampunk—with the 1990 alternate history novel <cite>The Difference Engine</cite>, written with Bruce Sterling. In the 1990s, he composed the Bridge trilogy of novels, which focused on sociological observations of near-future urban environments and late capitalism. His most recent novels—<cite>Pattern Recognition</cite> (2003), <cite>Spook Country</cite> (2007) and <cite>Zero History</cite> (2010)—are set in a contemporary world and have put his work onto mainstream bestseller lists for the first time.
<p>
Gibson is one of the best-known North American science fiction writers, fêted by <cite>The Guardian</cite> in 1999 as "probably the most important novelist of the past two decades". Gibson has written more than twenty short stories and ten critically acclaimed novels (one in collaboration), and has contributed articles to several major publications and collaborated extensively with performance artists, filmmakers and musicians. His thought has been cited as an influence on science fiction authors, design, academia, cyberculture, and technology.`,
	},
	{
		"William Shakespeare",
		"william_shakespeare.jpg",
		"", // Public domain, 1610
		"https://en.wikipedia.org/wiki/William_Shakespeare",
		`William Shakespeare (26 April 1564 (baptised) – 23 April 1616) was an English poet and playwright, widely regarded as the greatest writer in the English language and the world's pre-eminent dramatist. He is often called England's national poet and the "Bard of Avon". His extant works, including some collaborations, consist of about 38 plays, 154 sonnets, two long narrative poems, two epitaphs on a man named John Combe, one epitaph on Elias James, and several other poems. His plays have been translated into every major living language and are performed more often than those of any other playwright.
<p>
Shakespeare was born and brought up in Stratford-upon-Avon. At the age of 18, he married Anne Hathaway, with whom he had three children: Susanna, and twins Hamnet and Judith. Between 1585 and 1592, he began a successful career in London as an actor, writer, and part owner of a playing company called the Lord Chamberlain's Men, later known as the King's Men. He appears to have retired to Stratford around 1613 at age 49, where he died three years later. Few records of Shakespeare's private life survive, and there has been considerable speculation about such matters as his physical appearance, sexuality, religious beliefs, and whether the works attributed to him were written by others.
<p>
Shakespeare produced most of his known work between 1589 and 1613. His early plays were mainly comedies and histories, genres he raised to the peak of sophistication and artistry by the end of the 16th century. He then wrote mainly tragedies until about 1608, including <cite>Hamlet</cite>, <cite>King Lear</cite>, <cite>Othello</cite>, and <cite>Macbeth</cite>, considered some of the finest works in the English language. In his last phase, he wrote tragicomedies, also known as romances, and collaborated with other playwrights.`,
	},
}
