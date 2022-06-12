What is a telegram robot and what is its use?

Telegram bots are actually telegram accounts that are controlled and managed by software.
 These robots are able to perform any operation in this space such as training, ‌ play, search, play, ‌ connect, integrate with other services, restrict message sending in group..
As a result, you control your bots via the https protocol to the Telegram API. Telegram Robot API is also an HTTPS-based interface created for those interested in building robots in Telegram.
 Telegram bots are user accounts that do not require a phone number to work. These bots contain code that is executed on the server through a user interface, and in general, how these bots work is hidden from the user, and Telegram manages it on its own encrypted MTProto protocol.
In this way, the Telegram broker server controls all encrypted communications through an interface based on the HTTPS protocol, which is the Bot API.
Robots have limited memory in the cloud, meaning that old messages are deleted from the main server after they are fully processed. When you add a bot to a group, it cannot receive all the group messages unless it violates the Privacy Mode rules.
How much access does your Telegram robot have to your information?
The telegram rbot has access to a maximum of one normal user. So a telegram robot can not spy on your phone, nor can it access your passwords, nor can it send your gallery photos to anyone. . If you refer to the list of members of the Telegram group, the access level of the robot is written next to its name.

Build a simple telegram robot step by step with explanation

 interactive Telegram (responsive)

Set up Your Bot :

Go to the telegram app
Search for the “botfather” telegram bot.
Click on or type /newbot to create a new bot.
Follow instructions and make a new name for your bot. 
If you are making a bot just for experimentation, it can be useful to namespace your bot by placing your name before it in its username, 
since it has to be a unique name. Although, its screen name can be whatever you like.
I have chosen “arghvnBot” as the screen name and “arghvn_bot” as its username.
Congratulations! You have created your first bot. You should see a new API token generated for it (in .env file )
Now you can search for your newly created bot on telegram .

If we want to continue and chat with our robot,
The robot can not answer or say anything.
  Let's start by building our own robot server that runs in the back.





































What is the MTProto protocol?
The proprietary Telegram proxy, named MTProto, is not really just a new mechanism for bypassing Telegram filtering. This protocol (MTProto) was developed by Nikolai Dorf (Paul Dorf's brother).
A protocol is a set of rules and regulations that apply to communication.
The MTProto protocol stands for Mobile Telegram Protocol. A new and proprietary security protocol for Telegram that was developed and published by the same team and has been released in two versions to date (MTProto 1.0 and MTProto 2.0).
 Task: Transmits and exchanges encrypted messages from user to user (Client to Client) as End-To-End.
End-to-end or E2EE encryption lets you ensure that your data is kept secret until it reaches the intended recipient, meaning that no one else can see your private data. In other words, end-to-end encryption of the chat program is such that only you and the person you are talking to can read the content of your messages. In this scenario, even the company running the chat program will not be able to see the content of your messages.
For example, when you access your online banking website or other websites that use HTTPS, the communication between you and that website is encrypted to the network operator, ISP or any other person looking for your traffic. Is unable to view your bank password and financial details.
When connected to this proxy, messages and all exchanged data can only be decrypted and viewed by the sender and receiver (simple definition of End-To-End). In fact, this protocol (MTProto proxy) is not only provided for filtering, but also has a security enhancement aspect.

What happens when you are not connected to the Telegram proxy?
The cryptographic system mentioned above was for when you were connected to this proxy, now what difference would it make if a user referred to Telegram normally?
In this case, the message is encrypted, transmitted by the sender user to the Telegram servers and decrypted, and then re-encrypted and sent to the recipient user, in fact, it will be possible to listen to the messages on the Telegram servers.
The MTProto protocol can be related to the HTTPS protocol in some ways, but with the difference that the HTTPS protocol is an interface between the user and the server and only encrypts the information between the user (client) and the server (server) and the message is decrypted on the server and server Can see the content of the message.
Of course, this is not a drawback for this protocol, because the HTTPS protocol is basically designed and optimized for this purpose, namely: web browsing, file sharing, etc., and there are no restrictions on using this protocol for other purposes.
In addition to encrypting the content of messages, the MTProto protocol enables the structure of chats and metadata in encrypted form, while HTTPS does not. In fact, this protocol was developed solely for this application.

botfather intro :
he’s the one that’ll assist you with creating and managing your bot.
Botfather is a robot that can be used to build other robots. Working with this reference robot is very simple and almost everything you need to build a telegram robot is available in this robot.
To work with the Bot Fader robot, just search for the username @ Botfather in the Telegram search bar and then select it.
In the chat page with the robot, to start working, you must select the Start option to enter the starting stage to build the Telegram robot.

what is webhook ?
Webhook (Webhook) is an automated message that is sent by applications in the event of an event and is used in web programming.
 The webhook contains a payload packet that is sent to a unique URL.
Webhook works very similar to SMS notification.
Web hook is one of the ways in which web applications can communicate with each other. Third-party APIs can be used in the application development project using Web Hook. Web hook allows you to send real-time data from one application to another whenever an event occurs.
What is an event in Web Hook?
Events are human-triggered actions in an application. For example, clicking Add to Cart, sending a message to someone else, or heading to a specific webpage are all events triggered by people. Data related to these events.

In Pulling, we constantly request information from the Telegram server and receive the information.
But in the webhook discussion, we turn our local device into a server and it notifies us whenever the Telegram server receives a request.
Pulling puts more pressure on the system because it has to constantly check the Telegram server, but the webhook is more economical because it does not take action if there is no message.
But if the webhook is set up with one server, it can no longer be set up with another server, while Pooling can work with multiple servers at the same time.



















































Full description about building and launching a Telegram applicable robot

Set up Your Bot :
You don’t need to write any code for this.
 Go to the telegram app on your phone and Search for the “botfather” telegram bot 

 Type /help to see all possible commands the botfather can handle

 Click on or type /newbot to create a new bot.

 Follow instructions and make a new name for your bot. If you are making a bot just for experimentation, it can be useful to namespace your bot by placing your name before it in its username, since it has to be a unique name. Although, its screen name can be whatever you like.
I have chosen “arghvnbot” as the screen name and “arghvntelegrambot” as its username.
If your username has already been selected by someone, it will give an error.
The robot should send you the message "done ......" at the end.

Congratulations! You have created your first bot. You should see a new API token generated for it . 
Now you can search for your newly created bot on telegram :
searching username.

If we want to chat with our robot,
Well this is impossible.
Our robot can not answer or say anything.
We have to set up this possibility by building our robot server that runs in the back.

Set up Your Bot Server :

Every time you message a bot, it forwards your message in the form of an API call to a server. This server is what processes and responds to all the messages you send to the bot.
There are two ways we can go about receiving updates whenever someone sends messages to our bot :
1-Long polling : Periodically scan for any messages that may have appeared.
 Not recommended.
2-Webhooks : Have the bot call an API whenever it receives a message. 
Much faster and more responsive.

We are going to go with webhooks for this tutorial. Each webhook is called with an update object. Lets create our server to handle this update.

We will be creating our server using Golang, but you can use whatever suits you to make your server. Once you have Go installed, initialize your project:

initializing  our project  in GO in main.go

after coding in main.go , You can run this server on your local machine by running go run main.go
If you don’t see any error message, then that means your server is running on port 3000.
But, this is not enough. The bot cannot call an API if it is running on your local machine. It needs a public domain name. This means we have to deploy(Connect the tools) our application.

after ngrok Here and in the terminal an address like https: //e54851fb.ngrok.io is assigned to us.
Now, all we need to do is let telegram know that our bot has to talk to this url whenever it receives any message. We do this through the telegram API. Enter this in your terminal :

curl -F "url=https://e54851fb.ngrok.io/"  https://api.telegram.org/bot<your_api_token>/setWebhook

and you’re pretty much done! Try chatting with your newly made bot and see what happens!


About the BoConfigs fields :

1. APIKey: This is the token you received from botfather when you created the bot.

2. BotAPI: This is the address of the bot API server. If you are not running a local api server, use cfg.DefaultBotApi constant for this field.

3. UpdateConfigs: This is the update related configs such as update frequency (default : 300 milliseconds). You can use cfg.DefaultUpdateConfigs() method for this field.

3. Webhook: Pass true if you want to use webhook for receiving updates. (Telego currently does not support webhook).

4. LogFileAddress: The logs of the bot will be saved in this file.