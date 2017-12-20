package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.InvisibleAction;

import java.util.Map;
import java.util.logging.Logger;

/**
 * Class describing action performed on build page.
 */
@Extension
public class BuildReportAction extends InvisibleAction {

    public BuildReportAction(){}

    private String msg;
    private static final Logger LOGGER = Logger.getLogger(BuildReportAction.class.getName());

    public BuildReportAction(Map<String, String> report) {
        StringBuilder stb = new StringBuilder();
        for (Map.Entry<String,String> entry : report.entrySet()) {
            stb.append(entry.getKey());
            stb.append("\n\n");
            stb.append(entry.getValue());
            stb.append("\n\n");
        }
        msg = stb.toString();
        LOGGER.info("BuildReportAction running");
    }

    public String getMsg() {
        return msg;
    }
}
