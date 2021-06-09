import * as React from 'react';
import { Firehose, FirehoseResource } from '@console/internal/components/utils';
import { PodRingResources, PodRingData } from '../../types';
import { transformPodRingData, podRingFirehoseProps } from '../../utils';
import {
  PodModel,
  ReplicaSetModel,
  ReplicationControllerModel,
  DeploymentModel,
  DeploymentConfigModel,
} from '@console/internal/models';

interface RenderPropsType {
  loaded: boolean;
  loadError: any;
  data: PodRingData;
}

interface ControllerProps {
  loaded?: boolean;
  loadError?: any;
  resources?: PodRingResources;
  kind: string;
  render(RenderProps: RenderPropsType): React.ReactElement;
}

interface PodRingDataControllerProps {
  namespace: string;
  kind: string;
  render(RenderProps: RenderPropsType): React.ReactElement;
}

const Controller: React.FC<ControllerProps> = React.memo(
  ({ resources, render, loaded, loadError, kind }) => {
    return render({
      loaded,
      loadError,
      data: loaded ? transformPodRingData(resources, kind) : null,
    });
  },
);

const PodRingController: React.FC<PodRingDataControllerProps> = ({ namespace, kind, render }) => {
  const resources: FirehoseResource[] = [
    {
      isList: true,
      kind,
      namespace,
      prop: podRingFirehoseProps[kind],
    },
    {
      isList: true,
      kind: PodModel.kind,
      namespace,
      prop: podRingFirehoseProps[PodModel.kind],
    },
  ];

  switch (kind) {
    case DeploymentModel.kind:
      resources.push({
        isList: true,
        kind: ReplicaSetModel.kind,
        namespace,
        prop: podRingFirehoseProps[ReplicaSetModel.kind],
      });
      break;
    case DeploymentConfigModel.kind:
      resources.push({
        isList: true,
        kind: ReplicationControllerModel.kind,
        namespace,
        prop: podRingFirehoseProps[ReplicationControllerModel.kind],
      });
      break;
    default:
      break;
  }

  return (
    <Firehose resources={resources}>
      <Controller render={render} kind={kind} />
    </Firehose>
  );
};

export default React.memo(PodRingController);
