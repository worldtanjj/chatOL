var socket
$(document).ready(function () {
    //create socket
    socket = new WebSocket("ws://" + window.location.host + "/ws/join?uname=" + $('#uname').text());
    socket.onmessage = function (e) {
        console.log(e);
        var data = JSON.parse(e.data);
        console.log(data);
        switch (data.Type) {
            case 0: //join
                if (data.User == $('#uname').text()) {
                    $('#chatbox li').first().before("<li>You joined the chat room.</li>");
                } else {
                    $("#chatbox li").first().before("<li>" + data.User + " joined the chat room.</li>");
                }
                break;
            case 1: //leave
                $("#chatbox li").first().before("<li>" + data.User + " left the chat room.</li>");
                break;
            case 2: //message
                $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
                break;
        }
    }
    $("#sendbtn").click(function(){
        var uname = $('#uname').text();
        var content = $('#sendbox').val();
        socket.send(content);
        $('#sendbox').val("");
    })
})