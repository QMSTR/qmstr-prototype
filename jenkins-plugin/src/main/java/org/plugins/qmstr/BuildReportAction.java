package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.InvisibleAction;

import java.util.logging.Logger;

/**
 * Class describing action performed on build page.
 */
@Extension
public class BuildReportAction extends InvisibleAction {


    private String msg;
    private static final Logger LOGGER = Logger.getLogger(BuildReportAction.class.getName());

    public BuildReportAction() {
        this.msg = "Hello you";
        LOGGER.info("BuildReportAction running");
    }

    public String getMsg() {
        return msg;
    }
}
