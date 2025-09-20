package com.github.pratyayganguli;


import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {
    public static void main(String[] args) {
        Logger logger = LoggerFactory.getLogger(Main.class);
        var preExecutionTimeMillis = System.currentTimeMillis();
        logger.info("docker image for data has been successfully build");
        var postExecutionTimeMillis = System.currentTimeMillis();
        if (postExecutionTimeMillis > preExecutionTimeMillis) {
            logger.info("execution time taken: {} ms", postExecutionTimeMillis-preExecutionTimeMillis);
        } else {
            logger.error("quitting container");
        }
    }
}
