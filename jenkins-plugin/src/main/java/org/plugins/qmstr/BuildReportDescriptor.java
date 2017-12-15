package org.plugins.qmstr;

import hudson.model.AbstractProject;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.Publisher;

/**
 * Plugin connection to jenkins build page.
 */
public class BuildReportDescriptor extends BuildStepDescriptor<Publisher>{

    //TODO: Call the report

    /**
     * Get plugin availability.
     * @param jobType
     * 		Type of Job to we want this plugin to apply
     */
    public boolean isApplicable(
            final Class<? extends AbstractProject> jobType) {
        return true;
    }

    /**
     * Get Name of the plugin.
     */
    @Override
    public String getDisplayName() {
        return "Publish Reuse Report";
    }
}
