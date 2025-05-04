# Introdução

Em uma plataforma de streaming ao vivo como youtube, twitch, meta, tiktok existe um protocolo chamado RTMP que eu já cheguei a falar dele em outros vídeos, você pode clicar aqui para ver. Cheguei a fazer um vídeo aqui no canal no qual criamos um servidor de live streaming com NGINX simples, sem autenticação e nenhum outro requisito.

Agora vamos criar um servidor de live streaming com uma camada a mais de segurança: a chave de transmissão!

Por ser um projeto longo vamos dividir ele em 3 partes: A configuração inicial do servidor de live streaming e a nova diretiva on_publish, a segunda parte será o servidor de autenticação para que essa live aconteça, para isso vamos utilizar Go e a última parte exibindo o streaming para assistir também vamos criar um servidor para isso.


