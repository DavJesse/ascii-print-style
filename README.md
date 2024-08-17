<h1>ascii-art-web-export-file</h1>

<h2>Description</h2>
<p>Ascii-Art-Web-Export-File is now available as a web application with downloadable art! The web version of the tool generates ASCII art banners based on user-provided text input directly on the website, instead of using the command line. It continues to use predefined banner templates to create visually appealing representations of user input text.</p>

<h2>Features</h2>
<ul>
    <li>Generates ASCII art banners from input text</li>
    <li>Supports various banner templates, including "shadow," "standard," and "thinkertoy"</li>
    <li>Customizable text input</li>
    <li>Downloadable content with read/write privilages</li>
    <li>Responsive web interface</li>
</ul>

<h2>Usage</h2>
<ol>
    <li>Navigate to the web application URL.</li>
    <li>Enter the text you want to convert into an ASCII banner in the provided input field.</li>
    <li>Select the desired banner style (default is "standard").</li>
    <li>Click the "Generate" button to see the ASCII art banner displayed on the webpage.</li>
    <li>Should you wish to export the art, click on the 'Download Art' button.</li>
</ol>

<h2>Usage</h2>
<ol>
    <li>Clone the repository to your local machine.</li>
    <pre>From Gitea:<br>git clone https://learn.zone01kisumu.ke/git/davodhiambo/ascii-art-web-export-file.git<br><br>From Github:<br>git clone https://github.com/DavJesse/ascii-print-style.git</pre>
    <li>Navigate to the project directory.</li>
    <pre>For Gitea<br>cd ascii-art-web-export-file<br><br>For Github<br>cd ascii-print-style</pre>
    <li>Start server:</li>
    <pre>go run .</pre>
    <li>Access server locally from your browser</li>
    <pre>http://localhost:8000</pre></li>
</ol>

<h2>Note</h2>
<p>Listening ports on main.go can be replaced with any port from 1024 to 49151</p>

<h2>Example Usage</h2>
<p>Visit the web application, input the text "Hello", and select the "standard" style to generate the following ASCII art:</p>
<pre>
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
</pre>

<h2>Dependencies</h2>
<p>This tool requires no external dependencies. It is a standalone executable written in Go for the backend, with a simple frontend implemented in HTML and CSS.</p>

<h2>Authors</h2>
<p><a href="https://learn.zone01kisumu.ke/git/johnotieno0">John Paul Nyunja</a></p>
<p>Apprentice Software Developer, Zone01 Kisumu.</p>
<p>___________________________________________________</p>
<p><a href="https://learn.zone01kisumu.ke/git/davodhiambo">David Jesse Odhiambo</a></p>
<p>Apprentice Software Developer, Zone01 Kisumu.</p>
<p>___________________________________________________</p>
<p><a href="https://learn.zone01kisumu.ke/git/somulo">Samuel Omulo</a></p>
<p>Apprentice Software Developer, Zone01 Kisumu.</p>


---