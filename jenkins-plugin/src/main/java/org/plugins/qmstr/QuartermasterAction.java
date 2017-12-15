package org.plugins.qmstr;

import hudson.Extension;
import hudson.FilePath;
import hudson.Launcher;
import hudson.model.*;
import org.kohsuke.stapler.DataBoundConstructor;

import net.sf.json.JSONObject;
import org.kohsuke.stapler.StaplerRequest;
import jenkins.model.ParameterizedJobMixIn;

import java.io.IOException;
import java.util.logging.Logger;

public class QuartermasterAction extends JobProperty<Job<?, ?>>  {

    Process qmstr_master;

    @DataBoundConstructor
    public QuartermasterAction(String param) {
        /**
         * Start qmstr-master
         */

        /*try {

            qmstr_master = Runtime.getRuntime().exec("./qmstr-master");
            LOGGER.info("Qmstr-master is running " );
        } catch (IOException e){
            LOGGER.info("Could not find qmstr-master executable");
        }*/

        try {
            qmstr_master = new ProcessBuilder("", param).start();
            LOGGER.info("Qmstr-master is running " );
        } catch (IOException e){
            LOGGER.info("Could not find qmstr-master executable");
        }
    }

    public String getName() {
        return qmstr_master.toString();
    }

    @Extension
    public static final class DescriptorImpl extends JobPropertyDescriptor {
        /**
         * Used to hide property configuration under checkbox,
         * as of not each job is running with Qmstr build environment
         */
        public static final String QMSTRQ_PROJECT_BLOCK_NAME = "quartermasterProject";

        public boolean isApplicable(Class<? extends Job> jobType) {
            return ParameterizedJobMixIn.ParameterizedJob.class.isAssignableFrom(jobType);
        }

        public String getDisplayName() {
            return "execute with Quartermaster";
        }

        @Override
        public JobProperty<?> newInstance(StaplerRequest req, JSONObject formData) throws FormException {
            QuartermasterAction tpp = req.bindJSON(
                    QuartermasterAction.class,
                    formData.getJSONObject(QMSTRQ_PROJECT_BLOCK_NAME)
            );

            if (tpp == null) {
                LOGGER.fine("Couldn't bind JSON");
                return null;
            }

            return tpp;
        }

    }
    private static final Logger LOGGER = Logger.getLogger(QuartermasterAction.class.getName());

    /**
     * Perform the publication of the reuse report
     * @param build
     * 		Build on which to apply publication
     * @param workspacePath
     *      Unused
     * @param launcher
     * 		Unused
     * @param listener
     * 		Unused
     * @throws IOException
     * 		In case of file IO mismatch
     * @throws InterruptedException
     * 		In case of interuption
     */
    public void perform(final Run<?, ?> build, final FilePath workspacePath,
                        final Launcher launcher, final TaskListener listener)
            throws InterruptedException, IOException {

        /**
         * Show or not show reuse compliance badge
         */
        BuildReportAction buildAction;
        buildAction = new BuildReportAction(build);
        build.addAction(buildAction);




    }

}

