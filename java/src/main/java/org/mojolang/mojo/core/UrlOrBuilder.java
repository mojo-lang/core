// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/url.proto

package org.mojolang.mojo.core;

public interface UrlOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.Url)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string scheme = 1;</code>
   * @return The scheme.
   */
  java.lang.String getScheme();
  /**
   * <code>string scheme = 1;</code>
   * @return The bytes for scheme.
   */
  com.google.protobuf.ByteString
      getSchemeBytes();

  /**
   * <code>.mojo.core.Url.Authority authority = 2;</code>
   * @return Whether the authority field is set.
   */
  boolean hasAuthority();
  /**
   * <code>.mojo.core.Url.Authority authority = 2;</code>
   * @return The authority.
   */
  org.mojolang.mojo.core.Url.Authority getAuthority();
  /**
   * <code>.mojo.core.Url.Authority authority = 2;</code>
   */
  org.mojolang.mojo.core.Url.AuthorityOrBuilder getAuthorityOrBuilder();

  /**
   * <code>string path = 3;</code>
   * @return The path.
   */
  java.lang.String getPath();
  /**
   * <code>string path = 3;</code>
   * @return The bytes for path.
   */
  com.google.protobuf.ByteString
      getPathBytes();

  /**
   * <code>.mojo.core.Url.Query query = 5;</code>
   * @return Whether the query field is set.
   */
  boolean hasQuery();
  /**
   * <code>.mojo.core.Url.Query query = 5;</code>
   * @return The query.
   */
  org.mojolang.mojo.core.Url.Query getQuery();
  /**
   * <code>.mojo.core.Url.Query query = 5;</code>
   */
  org.mojolang.mojo.core.Url.QueryOrBuilder getQueryOrBuilder();

  /**
   * <code>string fragment = 7;</code>
   * @return The fragment.
   */
  java.lang.String getFragment();
  /**
   * <code>string fragment = 7;</code>
   * @return The bytes for fragment.
   */
  com.google.protobuf.ByteString
      getFragmentBytes();
}
