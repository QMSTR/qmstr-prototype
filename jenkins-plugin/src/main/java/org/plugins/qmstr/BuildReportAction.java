package org.plugins.qmstr;

import hudson.model.Action;
import hudson.model.Run;

import java.io.Serializable;

/**
 * Class describing action performed on build page.
 */
public class BuildReportAction implements Action, Serializable {

    private Run<?, ?> build;
    private String result;
    private String report;

    /**
     * Constructor
     */
    public BuildReportAction(final Run<?, ?> build) {
        this.build = build;

        //TODO: CREATE THE REPORT and call it from descriptor

    }

    /**
     * The three functions getIconFileName,getDisplayName,getUrlName create a
     * link to a new page with url : http://{root}/job/{job name}/URL_NAME for
     * the page of the build.
     */
    public String getIconFileName() {
        return null;
    }

    /**
     * The three functions getIconFileName,getDisplayName,getUrlName create a
     * link to a new page with url : http://{root}/job/{job name}/URL_NAME for
     * the page of the build.
     */
    public String getDisplayName() {
        return null;
    }

    /**
     * The three functions getIconFileName,getDisplayName,getUrlName create a
     * link to a new page with url : http://{root}/job/{job name}/URL_NAME for
     * the page of the build.
     */
    public String getUrlName() {
        // return URL_NAME;
        return null;
    }


    /**
     * Get Result.
     */
    public String getResult() {
        return this.result;
    }

    /**
     * Get Previous result.
     */
    String getPreviousResult() {
        BuildReportAction previousAction = this.getPreviousAction();
        String previousResult = null;
        if (previousAction != null) {
            previousResult = previousAction.getResult();
        }
        return previousResult;
    }

    /**
     * Get Previous action.
     */
    BuildReportAction getPreviousAction() {
        Run<?, ?> previousBuild = this.build.getPreviousBuild();
        if (previousBuild != null) {
            return previousBuild.getAction(BuildReportAction.class);
        }
        return null;
    }
}
