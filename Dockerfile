FROM  node:latest
RUN sudo apt-get update
RUN sudo apt-get install golang
RUN sudo npm install -g @tendermint/starport@0.0.10
WORKDIR /FindingImposter
COPY . .
EXPOSE 8080
EXPOSE 1317
EXPOSE 12345
EXPOSE 26657
EXPOSE 26656
CMD ["starport", "serve"]