FROM mongo:4.4

# Add a volume for MongoDB data
VOLUME /data/db

# Set up a working directory
WORKDIR /usr/src/data

# Expose port 27017
EXPOSE 27017

# Start MongoDB service
CMD ["mongod"]