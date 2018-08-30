package com.nineleaps.dockerworkshop.controllers;

import com.nineleaps.dockerworkshop.models.HealthCheck;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HealthCareController {

    private static final String OK ="ok";

    @RequestMapping(path = "/health-check")
    public HealthCheck getHealthStatus(){
        return new HealthCheck(OK);
    }
}
