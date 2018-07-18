import gql from "graphql-tag";

export interface DirectoryCreateVariables {
  name: string;
  parentId?: string;
}
const mDirectoryCreate = gql`
  mutation DirectoryCreate($name: String!, $parentId: ID) {
    createContainer(input: { name: $name, parentId: $parentId }) {
      id
      name
      parent {
        id
      }
      pages {
        id
      }
    }
  }
`;
export default mDirectoryCreate;